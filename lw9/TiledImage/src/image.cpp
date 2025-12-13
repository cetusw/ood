#include "image.h"
#include <algorithm>
#include <fstream>
#include <sstream>
#include <stdexcept>

Image::Image(const Size size, const Color color)
	: m_size(size)
{
	if (size.width <= 0 || size.height <= 0)
	{
		throw std::out_of_range("Image dimensions must be positive");
	}

	m_tilesX = (size.width + Tile::SIZE - 1) / Tile::SIZE;
	m_tilesY = (size.height + Tile::SIZE - 1) / Tile::SIZE;

	const CoW singleTile{ Tile(color) };
	m_tiles.assign(m_tilesX * m_tilesY, singleTile);
}

Size Image::GetSize() const noexcept
{
	return m_size;
}

Color Image::GetPixel(const Point p) const noexcept
{
	if (!IsPointInSize(p, m_size))
	{
		return 0;
	}

	const int tx = p.x / Tile::SIZE;
	const int ty = p.y / Tile::SIZE;
	const int lx = p.x % Tile::SIZE;
	const int ly = p.y % Tile::SIZE;

	return m_tiles[ty * m_tilesX + tx]->GetPixel({ lx, ly });
}

void Image::SetPixel(const Point p, const Color color)
{
	if (!IsPointInSize(p, m_size))
	{
		return;
	}

	const int tx = p.x / Tile::SIZE;
	const int ty = p.y / Tile::SIZE;
	const int lx = p.x % Tile::SIZE;
	const int ly = p.y % Tile::SIZE;

	m_tiles[ty * m_tilesX + tx].Write()->SetPixel({ lx, ly }, color);
}

void Image::SavePPM(const std::string& filename) const
{
	std::ofstream out(filename);
	if (!out)
		return;

	out << "P3\n";
	out << m_size.width << " " << m_size.height << "\n";
	out << "255\n";

	for (int y = 0; y < m_size.height; ++y)
	{
		for (int x = 0; x < m_size.width; ++x)
		{
			const Color c = GetPixel({ x, y });
			out << static_cast<int>(GetR(c)) << " "
				<< static_cast<int>(GetG(c)) << " "
				<< static_cast<int>(GetB(c)) << " ";
		}
		out << "\n";
	}
}