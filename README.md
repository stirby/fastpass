# FastPass

(it's not like the disney thing)

FastPass or `fp` is a command line password manager focused on getting you logged in quickly.

It uses fuzzy searching and learns which passwords you retrieve most frequently.

By default it generates easy to remember passwords using human words.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Example](#example)
- [Install](#install)
- [Features](#features)
- [Smart searching](#smart-searching)
- [Generators](#generators)
  - [Human](#human)
  - [Hex](#hex)
  - [Base62](#base62)
- [Password caching](#password-caching)
- [Recommended Name Format](#recommended-name-format)
- [Autocompletion?](#autocompletion)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Example

```bash
$ fp p
similar: ammarb@pizzahut.com ammarb@promhub.com 
Copied password for ammarb@papajohns.com
```

You should take care in making sure the right password is matched.

## Install

```bash
go get -u github.com/s-kirby/fastpass/cmd/fp
```

## Features 

- Encryption
- Fuzzy searching
- Notes
- Ranking based on access frequency
- Master password and key file support
- Key generated from master password caching
- Change master password
- Passwords are streched with pbkdf2/sha256 using 65536 iterations
- Multiple password generation strategies



## Smart searching

fp uses both password frequency and levenshtein distance distance from search to retrieve the best entry.

For example:

```bash
$ fp ls reddit
------------
/tmp/test.db: 5 entries
------------
usr1@reddit.com      [activity:003 created:04/03/2017 06:54:31]
usr2@reddit.com      [activity:000 created:04/03/2017 06:54:34]
ammarb@papajohns.com [activity:003 created:04/03/2017 06:53:29]
ammarb@pizzahut.com  [activity:000 created:04/03/2017 06:53:22]
ammarb@promhub.com   [activity:000 created:04/03/2017 06:53:37]
$ fp reddit
similar: usr2@reddit.com ammarb@papajohns.com ammarb@pizzahut.com ammarb@promhub.com 
Copied password for usr1@reddit.com
```


## Generators

### Human

The human password generator uses the lists in passgen/world_list/ to generate passwords.

It uses the following format: `<Adjective><Adjective><Noun><Noun><random num [000, 1000)>`

It generates about 55 bits of entropy.

### Hex

Hex generates 16 random hex digits.

It generates 64 bits of entropy.

### Base62

Base62 generates 16 random base62 digits.

It generates 96 bits of entropy.

## Password caching

fp caches secrets after an open in `/dev/shm/fp-<username>.secret`


## Recommended Name Format

FastPass imports and recommends the following name format, all lowercase.

`[category/...]<username>@<url/service>`

## Autocompletion?

I've decided to not add bash autocompletion as of now as it could leak account names after the database is closed.