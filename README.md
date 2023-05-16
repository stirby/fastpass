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

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Example

```bash
$ fp p
similar: ammarb36@pornhub.com ammarb36@pornhut.com ammarb36@papajohns.com ammarb36@pizzahut.com ammarb36@paypal.com 
ammarb36@pornhub.com -> Copied!
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
- Password and key file support
- Multiple password generation strategies
- Password caching
- Change password


## Smart searching

fp uses both password frequency and leviathan distance from search to retrieve the best entry.

For example:

```bash
ammar @ nebula > ~
$ fp list reddit
------------
/home/ammar/fastpass.db: 2 entries
------------
usr1@reddit.com  [hits:1 created:04/02/2017 14:49:50]
usr2@reddit.com  [hits:3 created:04/02/2017 14:48:58] Notes: hola passwords
ammar @ nebula > ~
$ fp reddit
similar: usr1@reddit.com 
usr2@reddit.com -> Password Copied!
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