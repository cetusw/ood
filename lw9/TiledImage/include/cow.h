#pragma once
#include <cassert>
#include <memory>
#include <utility>

template <typename Value>
class CoW
{
	struct WriteProxy
	{
		explicit WriteProxy(Value* value) noexcept
			: valuePtr{ value }
		{
		}

		WriteProxy(const WriteProxy&) = delete;

		WriteProxy& operator=(const WriteProxy&) = delete;

		Value& operator*() const& = delete;

		[[nodiscard]] Value& operator*() const&& noexcept
		{
			return *valuePtr;
		}

		Value* operator->() const& = delete;

		Value* operator->() const&& noexcept
		{
			return valuePtr;
		}

	private:
		Value* valuePtr;
	};

public:
	CoW()
		: m_value(std::make_shared<Value>())
	{
	}

	explicit CoW(Value&& val)
		: m_value(std::make_shared<Value>(std::move(val)))
	{
	}

	explicit CoW(const Value& val)
		: m_value(std::make_shared<Value>(val))
	{
	}

	CoW(const CoW& other) = default;

	CoW& operator=(const CoW& other) = default;

	CoW(CoW&& other) = default;

	CoW& operator=(CoW&& other) = default;

	const Value& operator*() const noexcept
	{
		assert(m_value);
		return *m_value;
	}

	const Value* operator->() const noexcept
	{
		assert(m_value);
		return m_value.get();
	}

	template <typename ModifierFn>
	void Write(ModifierFn&& modify)
	{
		EnsureUnique();
		std::forward<ModifierFn>(modify)(*m_value);
	}

	WriteProxy Write() && = delete;

	[[nodiscard]] WriteProxy Write() &
	{
		EnsureUnique();
		return WriteProxy(m_value.get());
	}

	Value& WriteBad()
	{
		EnsureUnique();
		return *m_value;
	}

private:
	void EnsureUnique()
	{
		assert(m_value);

		if (m_value.use_count() > 1)
		{
			m_value = std::make_shared<Value>(*m_value);
		}
	}

	std::shared_ptr<Value> m_value;
};
