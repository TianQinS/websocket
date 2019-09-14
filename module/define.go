package module

type Module interface {
	// OnInit initialize the module's mqtt topic and the lock-free queue's capacity.
	OnInit(topic string, qSize uint64)
	OnDestroy()
	// Run the module in goroutine.
	Run(closeSig chan bool)
	GetTopic() string
	RegisterMgr(mgr *ModuleManager)
	// GetMgr return the module manager.
	GetMgr() *ModuleManager
	// RegisterGo register a function for the global Post.
	RegisterGo(id string, f interface{})
	GetFunc(f string) (interface{}, error)
	// CallGo call a register function of the global Post in any gorountines.
	CallGo(f interface{}, params ...interface{}) error
	// CallSpec call a register function of the Post in the main event tick.
	CallSpec(f interface{}, params ...interface{}) error
	// Register register a function for this module.
	Register(id string, f interface{})
	// Call call a register function of this module.
	Call(f interface{}, params ...interface{}) error
	CallWithCallback(f, cb interface{}, cbParams, params []interface{}) error
	// ExecuteEvent process function calls on this module itself.
	ExecuteEvent() uint64
}
