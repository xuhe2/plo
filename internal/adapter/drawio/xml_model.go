package drawio

import "encoding/xml"

// MxFile 表示 Draw.io 文件的根元素
type MxFile struct {
	XMLName  xml.Name  `xml:"mxfile"`
	Diagram  Diagram   `xml:"diagram"`
}

// Diagram 表示图表
type Diagram struct {
	XMLName     xml.Name    `xml:"diagram"`
	MxGraphModel MxGraphModel `xml:"mxGraphModel"`
}

// MxGraphModel 表示图形模型
type MxGraphModel struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Root    Root     `xml:"root"`
}

// Root 表示根容器
type Root struct {
	XMLName xml.Name `xml:"root"`
	MxCells []MxCell `xml:"mxCell"`
}

// MxCell 表示单元格（节点或边）
type MxCell struct {
	XMLName  xml.Name  `xml:"mxCell"`
	ID       string    `xml:"id,attr"`
	Value    string    `xml:"value,attr,omitempty"`
	Style    string    `xml:"style,attr,omitempty"`
	Vertex   string    `xml:"vertex,attr,omitempty"`
	Edge     string    `xml:"edge,attr,omitempty"`
	Parent   string    `xml:"parent,attr,omitempty"`
	Source   string    `xml:"source,attr,omitempty"`
	Target   string    `xml:"target,attr,omitempty"`
	Geometry *MxGeometry `xml:"mxGeometry,omitempty"`
}

// MxGeometry 表示几何信息
type MxGeometry struct {
	XMLName  xml.Name `xml:"mxGeometry"`
	X        string   `xml:"x,attr,omitempty"`
	Y        string   `xml:"y,attr,omitempty"`
	Width    string   `xml:"width,attr,omitempty"`
	Height   string   `xml:"height,attr,omitempty"`
	Relative string   `xml:"relative,attr,omitempty"`
}

// IsVertex 检查是否为顶点（节点）
func (c *MxCell) IsVertex() bool {
	return c.Vertex == "1"
}

// IsEdge 检查是否为边
func (c *MxCell) IsEdge() bool {
	return c.Edge == "1"
}
