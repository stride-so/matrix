# Design

This document describes the design of the JSON competency matrix. JSON isn't a
particularly 'structured' data format but the layout itself benefits from some
explanation.

## The Matrix

A competency matrix can be thought of as a table encoding multiple dimensions of
information.  It describes the _Levels_ or stages of growth in a particular type
of organization. It also describes the skills and themes required in order to be
considered comptent or 'good' at that level of the role.

Skills belong to a `theme` which is an extra layer of annotation that provides
grouping of behaviour.  Each of these themes are either a technical skill or an
essential skill.

A technical skill is the trade-craft of that particular role, while essential
skills are the 'soft' skills surrounding it. Technical and esential skills are
not directly encoded into the model; they can be inferred from the text or
'semantic' definition of the theme you employ.

There is no existing or structured taxonomy of words for these areas.

The table can be thought of as such

```
-----------------------------------------------------------------------
| theme   | skill   | level 1                 | level 2    | level n  |
-----------------------------------------------------------------------
| theme 1 | skill 1 | skill 1 performed well  | ... better | ... best |
| theme 1 | skill 2 | skill 2 performed well  | ... better | ... best |
| theme n | skill n | skill n performed well  | ... better | ... best |
-----------------------------------------------------------------------
```

The JSON representation transposes much of this information into a structure
that is easier for a machine to read and write as well as `lint` or sanity
check. The JSON here is represnted as a javascript Object.

```
let matrix = {
  // for future sanity
  version: 0,
  // author of this document
  author: {
    name: "your name",
    email: "your email address",
  },
  company: {
    name: "your company name",
    url: "your matrix reference url",
  },
  // tracks are for encoding the graph of roles in a way that allow them to
  // connect up. For example you may be able to switch from being a independent
  // contributor to a manager that would be represented as an additional edge.
  tracks: [
    {
      name: "Independent Contributor",
      levels: ["Level 0", "Level 1"],
    },
  ],
  // themes are the grouped and named collections of skills.
  themes: [
    {
      title: "Collaboration",
      skills: ["Teamwork", "Communication", "Feedback"],
    },
    {
      title: "Leadership",
      skills: ["Mentoring & Learning", "Influence"],
    },
  ],
  // levels represent the stages in a career. For example junior, intermediate,
  // senior, staff, etc. Each of these are elements of an array with names and
  // their own array of titles.
  levels: [
    {
      name: "Level 0",
      track: "Independent Contributor",
      // titles is an array of strings that can be used to represent the role.
      titles: ["Junior Engineer"],
      detail:
        "You are learning quickly, and your new knowledge is benefitting your team..."
      skills: [
        {
          name: "Teamwork",
          body: "Helps teammates when asked. Gives and takes credit when due.",
        },
        {
          name: "Communication",
          body: "Can clearly articulate thoughts during meetings in both written and verbal form.",
        },
        {
          name: "Feedback",
          body: "Understands constructive feedback...",
        },
        {
          name: "Mentoring & Learning",
          body: "Seeks out opportunities for mentorship and education to improve their craft. ",
        },
        {
          name: "Influence",
          body: "Receives direction on tasks including their designs and ...",
        },
      ],
    },
  ],
  // additional levels go here
};
```

