# CIRA

### CIRA = Construction = Initialization = Return Referenced Resource Allocation

Many C++ programmers know the RAII idiom, that is, Resource acquisition is initialization (RAII). In C++ RAII actually
means, that resource acquisition must succeed for initialization to complete. Thus, all init is handled by the class
constructor and all de-init by the class destructor. The latter is required due to manual memory management in C++ which
requires to free all acquired memory by the end of an object lifecycle.

In Golang, however, thanks to the runtime, there is little need for a class destructor, but the overall idea that
resource acquisition must succeed for initialization remains valid. CIRA takes the RAI idea one step further and
specifies a more precise meaning of construction. Namely, in order to construct a standard component, the actual
constructor does three steps:

1. Construction
2. Complete initialization & acquisition
3. Return Referenced Resource Allocation

RAI and CIRA essentially aim for the same, specifically correct & reliable object creation. RAII gives more loose and
language agnostic guidance and as such remains applicable to many compiled languages such as C/C++, ADA, VALA, RUST, and
SWIFT. CIRA, to the contrary, gives very strict, precise, and very language specific guidance largely in response to
address very specific idiosyncrasies in the Golang programming language.

#### Construction

Unlike other several popular object-oriented, statically-typed programming languages, Golang does not have the notion of
a class. Instead, structs are types and methods are defined as functions attached to that type. Therefore, instead of
creating a class, in Golang the same gets accomplished by creating a struct. Just like C/C++, a struct can be
constructed by value or by reference. The component standard model specifies that any component gets constructed as a
reference.

#### Complete initialization & acquisition

One particularity in Golang relates to undefined default values of fields that have no assigned value yet. For example,
an empty reference is always of type nil. That is defined. An empty int, however, may have value 0 even though no value
has be defined. As a result, some non-initialized types cause a crash through unexpected nil values while other
non-initialized types cause undefined control flow behavior because of unexpected values. Because some types have
defined default and others not, it is surprising that Golang even let non-initialized types compile and cause all sorts
of strange runtime behavior.

Because Golang allows non-initialized types compile, the standard component model requires every single field to be
declared in a separate internal "state" sub-struct and every single value must be initialized during init. Furthermore,
all dependencies must be declared in a separate "deps" sub-struct and either be fully initialized or, when passed as
reference, tested to be non-nil.

In addition, the standard component model requires the init process to ensure five guarantees:

1) All internal fields have a value defined or are non-nil
2) All external dependencies are non-nil
3) All resources have been acquired i.e. connection has been established etc
4) Internal state, external dependencies or resources have been verified to be non-nil
5) The complete init & verify init process completed without error

If and only if all five guarantees have been fulfilled, the init process has been completed.

#### Return Referenced Resource Allocation

RAII does not state whether to return a value or a reference and to some extent, the answer depends on the specific
context. However, in the specific context of the standard component model, each standard component must return a
reference to its fully initialized implementation. This is because all standard components are initialized by the
standard component manager, which makes by default the assumption that global state is managed through shared
dependencies. More details are explained in the standard component model document.

From the perspective of the component manager, it becomes self-evident that

1) Each standard component must complete with a proper init or exit with a precise error
2) Each standard component must return a reference to its implementation
3) Each component that modifies shared data must protect access with a mutex.

Because of these requirements, the CIRA principle mandates:

1. Construct the main struct
2. Complete initialization & acquisition
3. Return (Reference) to Resource Allocation

Because construction calls init and a (successful) init returns a reference to all initialized and allocated resources,
the name CIRA fully reflects this practice:

**Construction is Initialization is Return Referenced Resource Allocation**
