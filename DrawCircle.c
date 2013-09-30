// The circle-drawing algorithm I came up with 
// in my first tutorial of KXT304 Computer Graphics & Animation unit 
// at Univerity of Tasmania,
// lectured by Tony Gray.
// Features one floating-point division, and no trig.

/*
 * DrawCircle
 * Draws a circle, the set of points (x,y), satisfying 
 *   (x-h)^2 + (y-k)^2 = r^2
 * for given parameters (h, k) (the centre) and r (radius).
 */
 
void DrawCircle(short h, short k, short r)
{
	float rx;
	short x,y;
	
	rx = r;
	x = r;
	y = 0;
	
	do 
	{
		plotPixel(h + x, k + y);
		plotPixel(h - x, k + y);
		plotPixel(h + x, k - y);
		plotPixel(h - x, k - y);
		plotPixel(h + y, k + x);
		plotPixel(h - y, k + x);
		plotPixel(h + y, k - x);
		plotPixel(h - y, k - x);
		
		rx -= (++y / rx);
		x = floor(rx + 0.5);
		
	} while (x >= y);
}