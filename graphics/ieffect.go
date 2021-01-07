package graphics

//IEffect is interface defining methods to add effects for various objects
type IEffect interface {
	//Apply applies effect to canvas
	Apply(canvas ICanvas)
}
