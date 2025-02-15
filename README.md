# The Operator Foundation

[Operator](https://operatorfoundation.org) makes usable tools to help people around the world with censorship, security, and privacy.

## Shapeshifter

The Shapeshifter project provides network protocol shapeshifting technology
(also sometimes referred to as obfuscation). The purpose of this technology is
to change the characteristics of network traffic so that it is not identified
and subsequently blocked by network filtering devices.

There are two components to Shapeshifter: transports and the dispatcher. Each
transport provide different approach to shapeshifting. These transports are
provided as a Go library which can be integrated directly into applications.
The dispatcher is a command line tool which provides a proxy that wraps the
transport library. It has several different proxy modes and can proxy both
TCP and UDP traffic.

If you are a tool developer working in the Go programming language, then you
probably want to use the transports library directly in your application.
<https://github.com/OperatorFoundation/shapeshifter-transports>

If you want a end user that is trying to circumvent filtering on your network or
you are a developer that wants to add pluggable transports to an existing tool
that is not written in the Go programming language, then you probably want the
dispatcher. Please note that familiarity with executing programs on the command
line is necessary to use this tool.
<https://github.com/OperatorFoundation/shapeshifter-dispatcher>

If you are looking for a complete, easy-to-use VPN that incorporates
shapeshifting technology and has a graphical user interface, consider
[Moonbounce](https://github.com/OperatorFoundation/Moonbounce), an application for macOS which incorporates shapeshifting without
the need to write code or use the command line.

### Shapeshifter Transports

This is the repository for the shapeshifter transports library for the Go
programming language. If you are looking for a tool which you can install and
use from the command line, take a look at the dispatcher instead:
<https://github.com/OperatorFoundation/shapeshifter-dispatcher>

The transports implement the Pluggable Transports 3.0 specification available here:
<https://github.com/Pluggable-Transports/Pluggable-Transports-spec/tree/main/releases/PTSpecV3.0> Specifically,
they implement the Go Transports API v2.1.

The purpose of the transport library is to provide a set of different
transports. Each transport implements a different method of shapeshifting
network traffic. The goal is for application traffic to be sent over the network
in a shapeshifted form that bypasses network filtering, allowing
the application to work on networks where it would otherwise be blocked or
heavily throttled.

The following transports are currently provided by the library:

#### Replicant

Replicant is Operator's flagship transport which can be tuned for each adversary. It is designed to be more effective and efficient that older transports.
It can be quickly reconfigured as filtering conditions change by updating just the configuration file.

You can find Replicant [here](https://github.com/OperatorFoundation/Replicant-go)

#### Optimizer

Optimizer is a pluggable transport that works with your other transports to find the best option. It has multiple configurable strategies to find
the optimal choice among the available transports. It can be used for numerous optimization tasks, such as round
robin load spreading among multiple transport servers or minimizing latency given multiple transport configurations.

You can find Optimizer [here](https://github.com/OperatorFoundation/Optimizer-go)

#### shadow (Shadowsocks)

Shadowsocks is a simple, but effective and popular network traffic obfuscation tool that uses basic encryption with a shared password.
shadow is a wrapper for Shadowsocks that makes it available as a Pluggable Transport.

You can find Shadow [here](https://github.com/OperatorFoundation/Shadow-go)

#### Installation

For individual installation instructions, see the README's for the transports:
[Shadow](https://github.com/OperatorFoundation/Shadow-go/blob/main/README.md)
[Replicant](https://github.com/OperatorFoundation/Replicant-go/blob/main/README.md)
[Optimizer](https://github.com/OperatorFoundation/Optimizer-go/blob/main/README.md)

The Shapeshifter transports are written in the Go programming language. To compile it you need
to install Go:

<https://golang.org/doc/install>

If you already have Go installed, make sure it is a compatible version:

    go version

The version should be 1.17 or higher.

If you get the error "go: command not found", then trying exiting your terminal
and starting a new one.

In order to use a transport in your project, you must have Go modules enabled in your project. How to do this is
beyond the scope of this document. You can find more information about Go modules here: <https://blog.golang.org/using-go-modules>

To use in your project, simply import the specific transport that you want to use, for example:

    import "github.com/OperatorFoundation/shapeshifter-transports/transports/Replicant/v2"
    
Your go build tools should automatically add this module to your go.mod and go.sum files. Otherwise, you can add it to the go.mod file directly. See the official Go modules guide for more information on this.    

Please note that the import path includes "/v3" to indicate that you want to use the version of the module compatible with the PT v3.0 specification. This is required by the Go modules guide.

When you build your project, it should automatically fetch the correct version of the transport module.

#### Frequently Asked Questions

##### What transport should I use in my application?

Try Replicant, Operator's flagship transport which can be tuned for each adversary. Email contact@operatorfoundation.org for a sample config file for the adversary of interest.
shadow is also a good choice as it works on many networks and is easy to configure.

If you are an application developer using Pluggable Transports, feel free to reach out to the Operator Foundation for
help in determining which transport might work best for your application. Email contact@operatorfoundation.org.

##### My application is not written in Go. Can I still use the transports?

Yes, the Go API is only one way to integrate transports into your application.
There is also an interprocess communication (IPC) protocol that allows you to
control a separate process (called the dispatcher) which provides access to the
transports through a proxy interface. When using this method, your application
can be written in any language. You just need to implement the IPC protocol so
that you can communicate with the dispatcher. The IPC protocol is specified in
the [Pluggable Transports 2.1 specification](https://www.pluggabletransports.info/spec/#build) section 3.3 and an implementation of the dispatcher is available which you can bundle with your
application here: <https://github.com/OperatorFoundation/shapeshifter-dispatcher>

### Credits

 * Operator Foundation - Replicant, Optimizer, and shadow (wrapper for Shadowsocks)
