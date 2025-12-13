#include "tile.h"

#include <cassert>

Tile::Tile(const Color color) noexcept
	: pixels()
{
	for (auto& row : pixels)
	{
		row.fill(color);
	}

	assert(m_instanceCount >= 0);
	++m_instanceCount;
}

Tile::Tile(const Tile& other)
	: pixels(other.pixels)
{
	assert(m_instanceCount >= 0);
	++m_instanceCount;
}

Tile::~Tile()
{
	--m_instanceCount;
	assert(m_instanceCount >= 0);
}

void Tile::SetPixel(const Point p, const Color color) noexcept
{
	if (IsPointInSize(p, { SIZE, SIZE }))
	{
		pixels[p.y][p.x] = color;
	}
}

Color Tile::GetPixel(const Point p) const noexcept
{
	if (IsPointInSize(p, { SIZE, SIZE }))
	{
		return pixels[p.y][p.x];
	}
	return 0;
}

int Tile::GetInstanceCount() noexcept
{
	return m_instanceCount;
}