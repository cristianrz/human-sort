# human-sort

Rank anything by answering pairwise comparisons.

Given a list of options, `human-sort` asks you to compare them two at a
time and produces a full ranking. It uses mergesort under the hood, so
the number of questions is O(n log n) — the minimum possible for a
comparison-based sort.

## Use cases

- Prioritising a backlog or to-do list
- Ranking job offers or apartments
- Ordering goals, ideas, or candidates
- Any list where the right order requires human judgement

## Install

```sh
go install github.com/cristianrz/human-sort@latest
```

## Usage

```sh
$ human-sort
Enter each option on a separate line. Press Ctrl+D when done.
Move to Berlin
Take the remote job
Go back to university
^D

[+] Which is better?
        1) Move to Berlin
        2) Take the remote job
> 2

[+] Which is better?
        1) Take the remote job
        2) Go back to university
> 1

...

Ranking:
  1. Take the remote job
  2. Move to Berlin
  3. Go back to university
```

## Why mergesort?

Mergesort divides the list in half recursively and only compares items
that need to be compared. For 10 items it asks at most 34 questions; a
naive approach would ask 45. For 20 items the difference is 69 vs 190.

## License

BSD 3-Clause
