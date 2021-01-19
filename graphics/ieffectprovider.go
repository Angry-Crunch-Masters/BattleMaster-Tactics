package graphics

//IEffectProvider is interface for providing effects from various sources
type IEffectProvider interface {
	ApplyEffects(canvas ICanvas)
}
