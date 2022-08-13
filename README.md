<div align="center">
<p>
    <img width="150" src="https://raw.githubusercontent.com/jule-lang/resources/main/jule.svg?sanitize=true">
</p>
<h2>The Jule Programming Language</h2>

[Website](https://jule-lang.github.io/website/) |
[Documentations](https://jule-lang.github.io/website/pages/docs.html) |
[Contributing](https://jule-lang.github.io/website/pages/contributing.html)

</strong>
</div>

<h2 id="motivation">Motivation</h2>

Jule is a statically typed compiled programming language designed for system development, building maintainable and reliable software.
The purpose of Jule is to keep functionality high while maintaining a simple form and readability.
It guarantees memory safety and does not contain undefined behavior.

<img src="./docs/images/quicksort.png"/>

<h2 id="memory-safety">Memory Safety and Management</h2>
The memory safety and the memory management.
A major challenge in the C and C++ or similar programming languages.
Jule guarantees memory safety and uses reference counting for memory management.
An account-allocation is automatically released as soon as the reference count reaches zero.
There are no dangling pointers, and accessing a null pointer will obviously get you an error.
<br><br>

+ Instant automatic memory initialization
+ Bounds checking
+ Null checking

<h2 id="cpp-interoperability">C++ Interoperability</h2>
Jule is designed to be interoperable with C++.
A C++ header file dependency can be added to the Jule code and its functions can be linked.
It's pretty easy to write C++ code that is compatible with the Jule code compiled by the compiler.
JuleC keeps all the C++ code it uses for Jule in its <a href="https://github.com/jule-lang/jule/tree/main/api">api</a> directory.
<ol></ol> <!-- for space -->
<img src="./docs/images/cpp_interop.png"/>

<h2 id="goals">Goals</h2>

+ Simplicity and maintainability
+ Fast and scalable development
+ Performance-critical software
+ Memory safety
+ As efficient and performance as C++
+ High C++ interoperability

<h2 id="what-is-julec">What is JuleC?</h2>
JuleC is the name of the reference compiler for the Jule programming language.
It is the original compiler of the Jule programming language.
The features that JuleC has is a representation of the official and must-have features of the Jule programming language.

<h2 id="about-project">About Project</h2>
JuleC, the reference compiler for Jule, is still in development.
Currently, it can only be built from source.
Due to the fact that it is still under development, there may be changes in the design and syntax of the language.
<br><br>
It is planned to rewrite the compiler with Jule after reference compiler reaches sufficient maturity.
JuleC has or is very close to many of the things Jule was intended to have, such as memory safety, properties, structures with methods and generics.
<br><br>
Currently, project structure, its lexical and syntactic structure has appeared.
However, when the reference compiler is rewritten with Jule, it is thought that AST, Lexer and some packages will be included in the standard library.
This will be a change that will cause the official compiler's project structure to be rebuilt.
The reference compiler will probably use the standard library a lot.
This will also allow developers to quickly develop tools for the language by leveraging Jule's standard library.

<h2 id="documentations">Documentations</h2>

All documentation about Jule is on the website. <br>
[See Documentations](https://jule-lang.github.io/website/pages/docs.html)

<h2 id="os-support">OS Support</h2>

<table>
    <tr>
        <td><strong>Operating System</strong></td>
        <td><strong>Transpiler</strong></td>
        <td><strong>Compiler</strong></td>
    </tr>
    <tr>
        <td>Windows</td>
        <td>Supports</td>
        <td>Not supports yet</td>
    </tr>
    <tr>
        <td>Linux</td>
        <td>Supports</td>
        <td>Not supports yet</td>
    </tr>
    <tr>
        <td>MacOS</td>
        <td>Supports</td>
        <td>Not supports yet</td>
    </tr>
</table>

<h2 id="building-project">Building Project</h2>

> [Website documentation](https://jule-lang.github.io/website/pages/docs.html?page=getting-started-install-from-source) for install from source.

There are scripts prepared for compiling of JuleC. <br>
These scripts are written to run from the home directory.

`build` scripts used for compile. <br>
`brun` scripts used for compile and execute if compiling is successful.

[Go to scripts directory](scripts)

JuleC aims to have a single main build file. <br>
JuleC is in development with the [Go](https://github.com/golang/go) programming language. <br>

### Building with Go Compiler

#### Windows - PowerShell
```
go build -o julec.exe -v cmd/julec/main.go
```

#### Linux - Bash
```
go build -o julec -v cmd/julec/main.go
```

Run the above command in your terminal, in the Jule project directory.

<h2 id="project-build-state">Project Build State</h2>

<table>
    <tr>
        <td><strong>Operating System</strong></td>
        <td><strong>State</strong></td>
    </tr>
    <tr>
        <td>Windows</td>
        <td>
            <a href="https://github.com/jule-lang/jule/actions/workflows/windows.yml">
                <img src="https://github.com/jule-lang/jule/actions/workflows/windows.yml/badge.svg")>
            </a>
        </td>
    </tr>
    <tr>
        <td>Ubuntu</td>
        <td>
            <a href="https://github.com/jule-lang/jule/actions/workflows/ubuntu.yml">
                <img src="https://github.com/jule-lang/jule/actions/workflows/ubuntu.yml/badge.svg")>
            </a>
        </td>
    </tr>
    <tr>
        <td>MacOS</td>
        <td>
            <a href="https://github.com/jule-lang/jule/actions/workflows/macos.yml">
                <img src="https://github.com/jule-lang/jule/actions/workflows/macos.yml/badge.svg")>
            </a>
        </td>
    </tr>
</table>

<h2 id="contributing">Contributing</h2>

Thanks for you want contributing to Jule!
<br><br>
The Jule project use issues for only bug reports and proposals. <br>
To contribute, please read the contribution guidelines from <a href="https://jule-lang.github.io/website/pages/contributing.html">here</a>. <br>
To discussions and questions, please use <a href="https://github.com/jule-lang/jule/discussions">discussions</a>.

<h2 id="code-of-conduct">Code of Conduct</h2>

[See Code of Conduct](https://jule-lang.github.io/website/pages/code_of_conduct.html)

<h2 id="license">License</h2>

The JuleC and standard library is distributed under the terms of the BSD 3-Clause license. <br>
[See License Details](https://jule-lang.github.io/website/pages/license.html)
