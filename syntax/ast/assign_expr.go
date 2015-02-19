package expr

import "encoding/json"

type BinaryExpr struct {
	SyntaxInfo
	Kind  string   `json:"kind"`
	Id    int      `json:"id"`
	Op    AssignOp `json:"op"`
	Left  *Expr    `json:"left"`
	Right *Expr    `json:"right"`
}

func (x *BinaryExpr) GetId() int {
	return x.Id
}

func (x *BinaryExpr) GetChildren() []Syntax {
	return []Syntax{Left, Right}
}

func (x *BinaryExpr) MarshalJSON() ([]byte, error) {

	return json.Marshal(*x)

}
func (x *BinaryExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
