package main

func main() {

}


type cell struct {
	x int
	y int
	cnt int
}

type 	grid struct {
	cells map[string]*cell
}

func (g *grid)setXy(x,y int) {
	if g.cells == nil {
		g.cells = map[string]cell{}
	}

	key := fmt.Sprintf("%d,%d",x,y)

	v,ok := g.cells[key]
	if !ok {
		v = &cell{x:x,y:y}
		g.cells[key]=v
	}
	v.cnt ++
}

func (g *grid)getVal(x,y int)int {
	key := fmt.Sprintf("%d,%d",x,y)

	v,ok := g.cells[key]
	if !ok {
		return 0
}
return v
}


func (g *grid)show(){
	maxx:=-9999
	minx:=9999

	maxy:=-9999
	miny:=9999


	for _, c:= range g.cells {
		if c.x <minx {
			miny = c.x
		}
		if c.x> maxx{
			maxy = c.x
		}
		if c.y <miny {
			miny = c.y
		}
		if c.y> maxy{
			maxy = c.y
		}
	}

	for idx,



}


type program struct {
	grd grid
	posX int
	posY int
}

func (p *program)run(input string) int{
	p.grd := grid{}
	steps := strings.Split(input)
	for _,step := range steps {
		dist,err := strconv.Atoi(step[1:])
		if err != nil {
			logrus.Fatal(err)
		}

		switch step[0] {
		case "R":
			for _, i := range dist {
				p.posX++
				p.grd.setXy(p.posX,p.posY)
			}
		case "L":
			for _, i := range dist {
				p.posX--
				p.grd.setXy(p.posX,p.posY)
			}
		case "D":
			for _, i := range dist {
				p.posY--
				p.grd.setXy(p.posX,p.posY)
			}
		case "U":
			for _, i := range dist {
				p.posY++
				p.grd.setXy(p.posX,p.posY)
			}
		}
	}

}
