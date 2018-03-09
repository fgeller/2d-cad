package main

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

var factor = float64(10.0)

type point struct {
	c    *gg.Context
	x, y float64
}

func (p *point) X() float64 { return factor * (10 + p.x) }
func (p *point) Y() float64 { return factor * (10 + p.y) }
func (p *point) lineTo(o *point) {
	p.c.DrawLine(p.X(), p.Y(), o.X(), o.Y())
}

func (p *point) yDimTo(o *point, left bool) {
	v := 1.0
	if left {
		v = -1.0
	}

	p.c.DrawLine(
		p.X(),
		p.Y(),
		p.X()+v*(5*factor),
		p.Y(),
	)
	p.c.DrawLine(
		p.X()+v*(5*factor),
		p.Y(),
		p.X()+v*(5*factor),
		o.Y(),
	)
	p.c.DrawLine(
		p.X(),
		o.Y(),
		p.X()+v*(5*factor),
		o.Y(),
	)

	s := fmt.Sprintf("%v", math.Abs(o.y-p.y))
	offset := 6.0
	if left {
		offset = 8.0
	}
	p.c.DrawString(s, p.X()+v*(offset*factor), p.Y()+(o.Y()-p.Y())/2)
}

func (p *point) xDimTo(o *point, onTop bool) {
	v := 1.0
	if onTop {
		v = -1.0
	}

	p.c.DrawLine(
		p.X(),
		p.Y(),
		p.X(),
		p.Y()+v*(5*factor),
	)
	p.c.DrawLine(
		p.X(),
		p.Y()+v*(5*factor),
		o.X(),
		o.Y()+v*(5*factor),
	)
	p.c.DrawLine(
		o.X(),
		p.Y(),
		o.X(),
		o.Y()+v*(5*factor),
	)

	s := fmt.Sprintf("%v", math.Abs(o.x-p.x))
	p.c.DrawString(s, p.X()+(o.X()-p.X())/2, p.Y()+v*(7*factor))
}

func drawPoly(ps ...*point) {
	for i := 0; i < len(ps)-1; i++ {
		ps[i].lineTo(ps[i+1])
	}
	ps[0].lineTo(ps[len(ps)-1])
}

func side() {
	var err error

	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(1)

	// seat
	sa := &point{dc, 10, 0}
	sb := &point{dc, 45, 0}
	sc := &point{dc, 45, 2.7}
	sd := &point{dc, 10, 2.7}
	drawPoly(sa, sb, sc, sd)

	// 35x10x4cm
	// top
	ta := &point{dc, 16.2, 0}
	tb := &point{dc, 45, 0}
	tc := &point{dc, 45, 3}
	tcd := &point{dc, 22, 7}
	td := &point{dc, 15, 7}
	drawPoly(ta, tb, tc, tcd, td)

	// 60x10x4cm -> 5cm offset, about 5degree angle
	la := &point{dc, 15, 7}
	lb := &point{dc, 22, 7}
	lc := &point{dc, 7, 67}
	ld := &point{dc, 0, 67}
	drawPoly(la, lb, lc, ld)

	lta := &point{dc, 15, 7 - 1.2}
	ltb := &point{dc, 22, 7 - 1.2}
	ltc := &point{dc, 22, 7}
	ltd := &point{dc, 15, 7}
	drawPoly(lta, ltb, ltc, ltd)

	lba := &point{dc, 0, 67}
	lbb := &point{dc, 7, 67}
	lbc := &point{dc, 7, 67 + 1.2}
	lbd := &point{dc, 0, 67 + 1.2}
	drawPoly(lba, lbb, lbc, lbd)

	// 40x10x4cm
	// base
	ba := &point{dc, 0, 67}
	bab := &point{dc, 7, 67}
	bb := &point{dc, 45, 72}
	bc := &point{dc, 45, 74}
	bd := &point{dc, -1.5, 74}
	drawPoly(ba, bab, bb, bc, bd)

	// 40x10x4cm
	// base support
	bsa := &point{dc, 40, 72}
	bsb := &point{dc, 45, 72}
	bsc := &point{dc, 45, 74}
	bsd := &point{dc, 40, 74}
	drawPoly(bsa, bsb, bsc, bsd)

	// 40x10x4cm
	// leg support
	bla := &point{dc, 2, 60}
	blb := &point{dc, 7, 60}
	blc := &point{dc, 7, 62}
	bld := &point{dc, 2, 62}
	drawPoly(bla, blb, blc, bld)

	dc.Stroke()

	dc.SetRGB(0, 1, 0)
	dc.SetLineWidth(1)
	sb.yDimTo(bc, false)
	sa.yDimTo(sd, true)
	ld.yDimTo(la, true)
	ba.yDimTo(bd, true)
	bd.xDimTo(bc, false)
	sa.xDimTo(sb, true)
	dc.Stroke()

	if err = dc.SavePNG("./stool-side.png"); err != nil {
		panic(fmt.Sprintf("failed to save png err=%v", err))
	}
}

func front() {
	var err error

	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(1)

	// seat
	sa := &point{dc, 2, 0}
	sb := &point{dc, 37, 0}
	sc := &point{dc, 37, 2.7}
	sd := &point{dc, 2, 2.7}
	drawPoly(sa, sb, sc, sd)

	// top left
	tlla := &point{dc, 0, 0}
	tllb := &point{dc, 2, 0}
	tllc := &point{dc, 2, 7}
	tlld := &point{dc, 0, 7}
	drawPoly(tlla, tllb, tllc, tlld)
	tlra := &point{dc, 2, 2.7}
	tlrb := &point{dc, 4, 2.7}
	tlrc := &point{dc, 4, 7}
	tlrd := &point{dc, 2, 7}
	drawPoly(tlra, tlrb, tlrc, tlrd)

	// top right
	trla := &point{dc, 35, 2.7}
	trlb := &point{dc, 37, 2.7}
	trlc := &point{dc, 37, 7}
	trld := &point{dc, 35, 7}
	drawPoly(trla, trlb, trlc, trld)
	trra := &point{dc, 37, 0}
	trrb := &point{dc, 39, 0}
	trrc := &point{dc, 39, 7}
	trrd := &point{dc, 37, 7}
	drawPoly(trra, trrb, trrc, trrd)

	// left leg
	llla := &point{dc, 0, 7}
	lllb := &point{dc, 2, 7}
	lllc := &point{dc, 2, 67}
	llld := &point{dc, 0, 67}
	drawPoly(llla, lllb, lllc, llld)
	llra := &point{dc, 2, 7}
	llrb := &point{dc, 4, 7}
	llrc := &point{dc, 4, 67}
	llrd := &point{dc, 2, 67}
	drawPoly(llra, llrb, llrc, llrd)

	// right leg
	rlla := &point{dc, 35, 7}
	rllb := &point{dc, 37, 7}
	rllc := &point{dc, 37, 67}
	rlld := &point{dc, 35, 67}
	drawPoly(rlla, rllb, rllc, rlld)
	rlra := &point{dc, 37, 7}
	rlrb := &point{dc, 39, 7}
	rlrc := &point{dc, 39, 67}
	rlrd := &point{dc, 37, 67}
	drawPoly(rlra, rlrb, rlrc, rlrd)

	// bottom left
	blla := &point{dc, 0, 67}
	bllb := &point{dc, 2, 67}
	bllc := &point{dc, 2, 74}
	blld := &point{dc, 0, 74}
	drawPoly(blla, bllb, bllc, blld)
	blra := &point{dc, 2, 67}
	blrb := &point{dc, 4, 67}
	blrc := &point{dc, 4, 74}
	blrd := &point{dc, 2, 74}
	drawPoly(blra, blrb, blrc, blrd)

	// bottom right
	brla := &point{dc, 35, 67}
	brlb := &point{dc, 37, 67}
	brlc := &point{dc, 37, 74}
	brld := &point{dc, 35, 74}
	drawPoly(brla, brlb, brlc, brld)
	brra := &point{dc, 37, 67}
	brrb := &point{dc, 39, 67}
	brrc := &point{dc, 39, 74}
	brrd := &point{dc, 37, 74}
	drawPoly(brra, brrb, brrc, brrd)

	// bottom support
	bsa := &point{dc, 4, 72}
	bsb := &point{dc, 35, 72}
	bsc := &point{dc, 35, 74}
	bsd := &point{dc, 4, 74}
	drawPoly(bsa, bsb, bsc, bsd)

	// leg support
	lsa := &point{dc, 4, 60}
	lsb := &point{dc, 35, 60}
	lsc := &point{dc, 35, 62}
	lsd := &point{dc, 4, 62}
	drawPoly(lsa, lsb, lsc, lsd)

	dc.Stroke()

	dc.SetRGB(0, 1, 0)
	dc.SetLineWidth(1)
	trrb.yDimTo(brrc, false)
	blld.xDimTo(brrc, false)
	tlla.xDimTo(tllb, true)
	dc.Stroke()

	if err = dc.SavePNG("./stool-front.png"); err != nil {
		panic(fmt.Sprintf("failed to save png err=%v", err))
	}
}

func main() {
	side()
	// front()
}
