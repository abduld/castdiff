package expr

import "encoding/json"

type BinaryExpr struct {
	SyntaxInfo
	Kind  string   `json:"kind"`
	Id    int      `json:"id"`
	Op    BinaryOp `json:"op"`
	Left  *Expr    `json:"left"`
	Right *Expr    `json:"right"`
}

func (x *BinaryExpr) GetId() int {
	return x.Id
}

func (x *BinaryExpr) String() string {
	str := x.Right.String()
	str += " " + x.Op.String() + " "
	str += x.Left.String()
	return str
}
func (x *BinaryExpr) GetChildren() []Syntax {
	return []Syntax{Left, Right}
}

func (x *BinaryExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "BinaryExpr"
	}
	return json.Marshal(*x)

}
func (x *BinaryExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
