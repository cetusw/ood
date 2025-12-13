#pragma once
#include "cow.h"
#include "geom.h"
#include "tile.h"
#include <string>
#include <vector>

class Image
{
public:
	explicit Image(Size size, Color color = 0x000000);

	[[nodiscard]] Size GetSize() const noexcept;

	[[nodiscard]] Color GetPixel(Point p) const noexcept;

	void SetPixel(Point p, Color color);

	void SavePPM(const std::string& filename) const;

private:
	Size m_size;
	int m_tilesX;
	int m_tilesY;
	std::vector<CoW<Tile>> m_tiles;
};