# FastPass

(it's not like the disney thing)

FastPass or `fp` is a command line password manager focused on getting you logged in quickly.

It uses fuzzy searching and learns which passwords you retrieve most frequently.

By default it generates easy to remember passwords using human words.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [FastPass](#fastpass)
  - [Example](#example)
  - [Install](#install)
  - [Features](#features)
  - [Generators](#generators)
    - [Human](#human)
    - [Hex](#hex)
    - [Base62](#base62)
  - [Password caching](#password-caching)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Example

```bash
$ fp p
other matches: ammarb36@pornhub.com ammarb36@pornhut.com ammarb36@papajohns.com ammarb36@pizzahut.com ammarb36@paypal.com 
ammarb36@pornhub.com -> Copied!
```

You should take care in making sure the right password is matched.

## Install

```bash
go get -u github.com/ammario/fastpass/cmd/fp
```

## Features 

- Encryption
- Fuzzy searching
- Notes
- Ranking based on access frequency
- Password and key file support
- Multiple password generation strategies
- Password caching
- Change password


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