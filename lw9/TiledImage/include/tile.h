#pragma once
#include "geom.h"
#include <array>

class Tile
{
public:
	constexpr static int SIZE = 8;

	explicit Tile(Color color = 0x000000) noexcept;

	Tile(const Tile& other);

	~Tile();

	void SetPixel(Point p, Color color) noexcept;

	[[nodiscard]] Color GetPixel(Point p) const noexcept;

	static int GetInstanceCount() noexcept;

private:
	inline static int m_instanceCount = 0;

	std::array<std::array<Color, SIZE>, SIZE> pixels;
};