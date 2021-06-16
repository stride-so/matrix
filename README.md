# Software Engineering Competency Matrix

This repository contains an ["Open Competency Matrix"][2] for Software
Engineers. It includes a standard data structure as well as a basic tool for
transposing that data structure into friendlier formats like job postings and
tables for human consumption.

This is an open-source project and all contents here are licensed appropriately.
You are encouraged to add your own organizations skills matrix to this
repository. You can do so by making a copy of the existing stride matrix model
and modifying it to suit your organization.

## Getting Started

```
$ git clone git@github.com/stride-so/matrix.git
```

## What is a Competency Matrix?

A skills competency matrix is a **tool** that is used to provide team members
with information on what a roles expectations are. By extension the matrix also
communications what skills and capabilities are required in order to advance
to the next stage of their career. While a competency matrix is not a
comprehensive list of every single skill, it is a framework that helps you have
conversations with your managers to get more actionable advice and information. 
 
If you or your team do not have a competency matrix in place, you should
encourage them to adopt one.

## How do I read a Competency Matrix?

This repository contains a `models` directory which contains JSON encoded
representations of a competency `matrix`.  Each matrix file encodes several
dimensions of information.

1. The levels and titles used within the matrix.
2. The title and body of any number of skills. Skills may be technical skills or
they may be essential skills.
3. Skills are grouped into areas called 'themes'

The JSON representation is a transposition and 'structured' encoding of a table
[like this one][2]. The linked table is parsed to generate the JSON file whose
format and layout is [described in this document][4]. 

## FAQ

### Can I add my company's matrix?

Yes! Please!

### How do I make my own matrix?

You can make your own matrix by copying what already exists and changing it
to something that suits your organization better. You can copy the [stride open
competency matrix][4] to something you can edit, then use that without
limitation. You can also copy the JSON representation in the models folder and
change your new file.

If you would like to include your matrix in this repository there is a tool that
lets you convert the matrix layout of the [example comptency matrix][2] into the
JSON representation. 

1. Copy the Google Sheet
2. make your edits then export the sheet as an `.xlsx` file. 
3. Then run the tool against that file. 

You will need `go` installed on your machine.

```shell
$ go install github.com/stride-so/matrix/tools/cmd/stride@latest
$ stride parse -in filename.xlsx > filename.json
```

### Why don't the role titles match my organization?

Titles mean a lot to some people, but titles are not the same across all
organizations. For this reason the titles in the provided model are an attempted
standardization across multiple organizations. This was done by collecting, and
reading through a list of about 20 different competency matrixes, career tracks,
etc from several software organizations. If you have friendly suggestions on
more appropriate titles please [open a pull request][3] or feel free to add your
own matrix.

### What is the difference between Essential and Technical Skills?

Essential skills are sometimes referred to as transferable or 'soft' skills.
Technical skills apply to a specific domain. For example, writing documentation,
being a project leader, understanding budgets, leadership, mentoring,
communication, etc are all essential skills. No matter which role you have at
any company you will likely need to employ some (or all!) of these skills
depending on your role. In a software capacity, technical skills includes
composition and comprehension of code as well as understanding testing,
security, hardware, constraints, trade-offs, etc.

<!--- links --->
[2]: https://docs.google.com/spreadsheets/d/1Qj2FTqKsuHPqnMyDNUFKi-RiiCU77EWuQwVODS4lS18/edit#gid=1772650517
[3]: https://github.com/stride-so/matrix/compare
[4]: ./docs/DESIGN.md
