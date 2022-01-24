# Standard Component Model

## Synopsis.

The OMX system follows the standard component model outlined in this document. The component model was initially
developed and refined throughout two previous golang projects and still keeps evolving. This document elaborates the
standard component model, its intent, and rational underlying its specific details. The intent of the model is to build
and share reliable, robust, and reusable code in a uniform and productive way. Uniform code structure really serves the
purpose of fast code navigation since each component is equally structured, the only focus really is the actual
implementation of its interface since everything else remains structurally constant.

Productivity results from a relatively straight simple and forward process to implement and integrate components. Since
the model establishes the convention that each component can only be crated and configured through a central component
manager, it essentially means that component integration really follows a uniform pattern too. Furthermore, productivity
also follows from template driven component generation that delivers a complete ready to use component pretty much from
the moment it has been requested through a shell script which means only interface definition and implementaiton has to
be actually done with all structures, files, and folders already in place.

## Principles

Three principles form the foundation of the standard component model:

1) Templates
2) Reliable, Robust, and Reusable
3) CIRA
4) Workflow
5) IoC

### Templates

The standard component model gets its name from its standardized file and folder structure. To ensure uniform file &
folder structures, component templates are versioned in a dedicated git repository, downloaded, and adapted via a shell
script. The exact details can be seen in the corresponding repository and shell script. In this document, the overall
structure and underlying intent will be elaborated.

### Reliable, Robust, and Reusable

The true intent of the model is to share reliable, robust, and reusable code. As a byproduct, cognitive load reduces
substantially during programming since each component follows a uniform structure and behaves in a uniform way.
Specifically:

1) Reliable means components work as defined
2) Robust means clear state handling & error handling
3) Reusable means dependencies are managed in a way that components can be re-used in a different context.

#### Reliable

Every standard component comes with an interface, a default constructor, and a default v1 implementation. The interface
defines the contract between the component and all external usage, the default constructor matches the interfaces to its
actual implementation and the default v1 implementation provides exactly that implementation. For example, the account
manage shown below has the following default interface & constructor:

```go
type AccountManager interface {
GetAccount(clientID, ticker string) (c *rest.Client, err error)
SetLeverage(clientID, ticker string, leverage int) (ok bool, msg string)
ResetLeverage(clientID, ticker string) (ok bool, msg string)
}

func NewAccountManager(apiManager api_manager.ApiManager)  AccountManager {
return v1.NewAccountManager(apiManager)
}
```

When for some reason, multiple different implementations become necessary, then a simple enum parameter can be used to
switch between implementations and return the requested one. This is very important for maintainability in the sense of
preserving the interface contract, which in turn makes the component more robust and reliable.

#### Robust

Robust software does not mean the absence of bugs, but rather continues functionality despite errors may occur. in
Golang, there are roughly three categories of potential errors than can occur in any non-trivial software

1) Initialization error i.e. wrong or nil value when it should not occur
2) External error i.e. connection is down when it is expected to be available.
3) Logic error i.e. actual behavior differs from specified expectation

The CIRA document discusses component initialization in detail and how exactly it prevents the bulk of standard
initialization erros. A large part of the resulting robustness comes from the separation of state and dependencies.
Conventionally, programs are scattered with variables with some being local to a function and others are shared between
functions. The standard model requires that every single variable that is shared between function must be declared
inside a dedicated state sub-struct which then gets fully initialized as described in the CIRA document. In the
aforementioned AccountManager, internal state requires a mutex and a custom hashmap.

```go
type State struct {
sync.RWMutex
clientMap *types.SyncedOrderedMap[string, *rest.Client]
}

func newState() (state *State) {
state = &State{
clientMap: types.NewSyncedOrderedMap[string, *rest.Client](),
}
return state
}
```

The purpose really is to centralize all shared fields in one single place and uniformly access them through the state
sub-struct. Furthermore, the state prefix states explicit that a non-local field is accessed.

Likewise, the dependencies' struct serves the same purpose of centralizing all dependencies in a central place and
providing a prefix explicitly stating a non-local method is accessed. For the AccountManager, just a single dependency
will be injected through the component constructor but the Dependencies struct still holds it. In more complex
components with many dependencies, for example the component managers, this comes very handy as it greatly simplifies
side by side inspections.

```go
type Dependencies struct {
apiManager api_manager.ApiManager
}

func newDependencies(apiManager api_manager.ApiManager) (deps *Dependencies) {
deps = &Dependencies{apiManager: apiManager}
return deps
```

In terms of robust & reliable components, the default constructor applies the CIRA principle by fully initializing all
state variables and dependencies and only returns a fully functional component. The AccountManager constructor is shown
below.

```go
type AccountManager struct {
dependencies *Dependencies
state        *State
}

func NewAccountManager(apiManager api_manager.ApiManager) *AccountManager {
comp := AccountManager{ // CIRA = Construction = Initialization = Return (Resource) Allocation
newDependencies(apiManager),
newState()} // 1. Construction
comp.init() // 2. Initialization
return &comp // 3. Return Reference to (Resource) Allocation
}
```

#### Reusable

Standard components only rely on public interfaces of other standard components, which makes them relatively easy to
replace with a mock for testing i.e. providing an in-memory mock implementation instead of full datawarehouse
integration while preserving the full interface.

## CIRA

The CIRA document elobarates in detail the complete initialization process of a standard component. For the purpose of
this document, it is enough to summarize CIRA as a standard mechanism to ensure each component only completes its
construction when its fully configured, functional, and verified.

## Workflow

With state and dependencies fully standardized across all components, the actual workflow sits at the core of each
standard component. Fundamentally, each workflow file groups one set of methods together i.e. all account methods for
the AccountManager. Because the account manager holds shared state, that is, data tha will be used across different
components, it must safeguard all read & write access via mutex. For example, the GetAccount methods follows the
standard Go practice to do so:

```go
func (c AccountManager) GetAccount(clientID, ticker string) (client *rest.Client, err error) {
c.state.RLock()
defer c.state.RUnlock()

client, exists := c.state.clientMap.Get(clientID + ticker)
if exists {
...
`	}
```

A pretty standard method only with the addition of a state struct to hold all fields used across methods and, not shown
in the sample, the dependency struct. Because state is centrally defined, the standard function in most IDE's "find
usage" helps to debug rare cases of incorrect order of value write. For the most part, mutex protection prevents those
issues in the first place, but very complex control flow in a method may cause hard to manage state changes. Speaking of
those complex cases, it might be favorable to implement them as state machines to ascertain explicitly every possible
state change.

## IoC - Inversion of Control

All standard components are initialized by the standard component manager, which applies the standard practice of
inversion of control. The standard component models refines IoC in the sense that global state is managed through shared
singleton components by default. Doing so solves the problem of centralized state management using only tools already
available instead of introducing additional and unnecessary complexity.  
The practice of managing shared state through a singleton components follows from the CIRA principle of only returning a
reference (pointer)
to the instance of the component.

To illustrate the practicality of the solution, let's consider the alternative of a component returning a value instead.
If a component would return a value, each instance would create a new and complete copy so that state is scattered
across multiple copies and actually really hard to manage because some kind of coordination would become required. Much
worse, if standard components would inject dependencies as values, then each receiving component would create and modify
it own local copy, which makes managing global state factually impossible.

### Properties

Therefore, returning a reference to the component preserves the following properties:

1) Singleton instantiation in the component manager. Each component exists once and only once.
2) Passed by reference. Each component takes its dependencies as reference only. This ensures that all components point
   to a single instance.
3) Simplified null check. Both, the component manager and each receiving component can easily check the validity of the
   pointer.

### Constraints

The singleton shared by reference model requires a few but very important constraints:

1) Each data read / write in a shared data manager component must be mutex protected.
2) Each connection component must be embedded exactly once into a data manger that protects each read & write with
   mutex, but preserves a single connection pool.
3) The data manger, on the other hand, can be embedded in as a many other components as the requirements demand b/c it
   ensures safe concurrent data access through mutex protection.

For example, a database client gets instantiated exactly once, then passed by reference into a data manager standard
component with mutex protected read / write functions to multiple DB tables, and then the data manager gets passes it by
reference into multiple other components essentially all of them holding pointers to just one singleton instance of the
data manger with just one DB client.

### Reuse

By design, the same component manager can easily be integrated into a web service, a commandline tool, or a desktop
application thus reusing some or all components in different application models and even across different applications.
For example, a shared component repository can be used across different applications as it would be the case of a
classic mono-repo. In fact, a mono-repo is the recommended structure for the standard component model exactly because of
the high degree of code reuse that follows when using both combined. For security reasons, it may be advisable to
re-implement the component manager with an application specific scope as to only provided components actually used in a
specific application instead to give access to all components of the repository.

### Con's

1) Adds quite some extra amount of boilerplate sourcecode. Other models might be more compact and code efficient.
2) Requires a disciplined, consistent, and uniform usage. May not work in teams with different coding styles.
3) Unfavorable overhead for a small code bases i.e. below ~15k LoC. The model excels at code bases of at least 50k Loc.

### Pro's

Despite the stated constraints and structural disadvantages, the standard component model also has some upsides:

1) Simple integration: Create a component manager and request any component. That's it for integration.
2) Generally low application memory consumption (~ 3-5 Mb) because of only references to singleton instances.
3) Generally safe data access because of required mutex protection in shared data components.

### Discussion

Obviously, there are certain use cases which necessitate a new instance of a component that may not hold any shared
state, but in any case this must be implemented in the component manager and documented accordingly i.e. through
appropriate method naming convention that reflects the fact. That does not invalidate the standard component model,
which prefers references, but rather acknowledges and embraces different requirements.

Thread-safety and safe parallel processing depends much more on the correct usage of Golang's language features i.e.
Goroutines than the actual component model that only encapsulates the code. Since components are assumed to be singleton
by default, managing internal state used across different threads or goroutines really comes down to the standard tools
and best practices.

Golang does not come with zero-cost abstraction as Rust does, which means there is a certain overhead to using
interfaces. In a previous project with low latency requirement a very similar system was used minus the interfaces. From
practical experiences, the main disadvantage comes from dependencies being linked to a specific implementation type
rather than a neutral type. When those dependencies rarely change, the performance gains in terms of reduces overhead
and lower latency clearly warrant the absence of interfaces. However, when dependencies require some variation, tend to
evolve, or require different implementation details depending on the actual context, then the interface approaches
outweigh most other concerns as long as raw performance isn't the deciding factor.

### Best practices

From the experience of using the standard component model in 2 -3 Golang only projects, it became self-evident that
several best practice make it so much less of an overhead and actually quite fun to use. Specifically:

1) Organize a mono-repo with a shared component pool.
2) Standardize everything you build more than once i.e. microservices and make it a versioned template
3) Automate templating with bash scripts as much as possible to generate components & build files
4) Build the mono-repo with bazel and generate build files with Gazelle
5) Abstract complexity away with a comprehensive makefile that covers most daily dev tasks 

