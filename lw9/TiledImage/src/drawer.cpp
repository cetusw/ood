#include "drawer.h"
#include <cmath>

namespace
{

int GetSign(const int value)
{
	return (0 < value) - (value < 0);
}

void DrawShallowLine(Image& image, const Point from, const Point to, const Color color)
{
	const int dx = std::abs(to.x - from.x);
	const int dy = std::abs(to.y - from.y);
	const int stepY = GetSign(to.y - from.y);
	int error = dx / 2;
	int y = from.y;

	for (int x = from.x; x <= to.x; ++x)
	{
		image.SetPixel({ x, y }, color);
		error -= dy;
		if (error < 0)
		{
			y += stepY;
			error += dx;
		}
	}
}

void DrawSteepLine(Image& image, const Point from, const Point to, const Color color)
{
	const int dx = std::abs(to.x - from.x);
	const int dy = std::abs(to.y - from.y);
	const int stepX = GetSign(to.x - from.x);
	int error = dy / 2;
	int x = from.x;

	for (int y = from.y; y <= to.y; ++y)
	{
		image.SetPixel({ x, y }, color);
		error -= dx;
		if (error < 0)
		{
			x += stepX;
			error += dy;
		}
	}
}

void PlotCirclePoints(Image& image, const Point center, const int x, const int y, const Color color)
{
	image.SetPixel({ center.x + x, center.y + y }, color);
	image.SetPixel({ center.x - x, center.y + y }, color);
	image.SetPixel({ center.x + x, center.y - y }, color);
	image.SetPixel({ center.x - x, center.y - y }, color);
	image.SetPixel({ center.x + y, center.y + x }, color);
	image.SetPixel({ center.x - y, center.y + x }, color);
	image.SetPixel({ center.x + y, center.y - x }, color);
	image.SetPixel({ center.x - y, center.y - x }, color);
}

void DrawCircleFillLines(Image& image, const Point center, const int x, const int y, const Color color)
{
	DrawLine(image, { center.x - x, center.y + y }, { center.x + x, center.y + y }, color);
	DrawLine(image, { center.x - x, center.y - y }, { center.x + x, center.y - y }, color);

	DrawLine(image, { center.x - y, center.y + x }, { center.x + y, center.y + x }, color);
	DrawLine(image, { center.x - y, center.y - x }, { center.x + y, center.y - x }, color);
}

} // namespace

void DrawLine(Image& image, Point from, Point to, const Color color)
{
	const int deltaX = std::abs(to.x - from.x);
	const int deltaY = std::abs(to.y - from.y);

	if (deltaY > deltaX)
	{
		if (from.y > to.y)
		{
			std::swap(from, to);
		}
		DrawSteepLine(image, from, to, color);
	}
	else
	{
		if (from.x > to.x)
		{
			std::swap(from, to);
		}
		DrawShallowLine(image, from, to, color);
	}
}

void DrawCircle(Image& image, const Point center, const int radius, const Color color)
{
	int x = 0;
	int y = radius;
	int delta = 3 - 2 * radius;

	while (y >= x)
	{
		PlotCirclePoints(image, center, x, y, color);
		x++;
		if (delta > 0)
		{
			y--;
			delta += 4 * (x - y) + 10;
		}
		else
		{
			delta += 4 * x + 6;
		}
	}
}

void FillCircle(Image& image, const Point center, const int radius, const Color color)
{
	int x = 0;
	int y = radius;
	int delta = 3 - 2 * radius;

	while (y >= x)
	{
		DrawCircleFillLines(image, center, x, y, color);
		x++;
		if (delta > 0)
		{
			y--;
			delta += 4 * (x - y) + 10;
		}
		else
		{
			delta += 4 * x + 6;
		}
	}
}