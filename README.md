# Nottu

Simple notepad in our browser

> nottu ui is still in developement

## Introduction

To use or contribute to nottu , you need to install Go first in your system and set its workspace

```sh
$ go get -u github.com/5anthosh/nottu
```

To run nottu

```
$ nottu
```

---

## Create

Used to create a new note

**URL** : `/notes`

**Method** : `POST`

**Auth required** : NO

### Data constraints

```json
{
  "title": "[title]",
  "content": "[content 1]"
}
```

### Data example

```json
{
  "title": "note123",
  "content": "content 1"
}
```

### Success Response

**Code** : `201 Created`

### Content example

```json
{
  "data": {
    "id": "bla6fss5p58afdc8tt7g",
    "message": "Note is created successfully"
  }
}
```

---

## Get

Used to get the all notes

**URL** : `/notes`

**Method** : `GET`

**Auth required** : NO

**Data constraints required** : NO

### Success Response

**Code** : `200 OK`

### Content example

```json
{
  "data": {
    "notes": [
      {
        "id": "bkpg9ac5p58bs8eoh9ig",
        "title": "note1",
        "content": "content 1",
        "created": "2019-07-20T17:40:17.81487838+05:30"
      }
    ]
  }
}
```

---

## Get by ID

Used to get the a note by its ID

**URL** : `/notes/:id`

**Method** : `GET`

**Auth required** : NO

**Data constraints required** : NO

### Success Response

**Code** : `200 OK`

### Content example

```json
{
  "data": {
    "notes": [
      {
        "id": "bkpg9ac5p58bs8eoh9ig",
        "title": "note1",
        "content": "content 1",
        "created": "2019-07-20T17:40:17.81487838+05:30"
      }
    ]
  }
}
```

---

## Update

Used to update the note

**URL** : `/notes/:id`

**Method** : `PUT`

**Auth required** : NO

### Data constraints

_Update both title and content_

```json
{
  "title": "[title]",
  "content": "[content 1]"
}
```

_Update only title_

```json
{
  "title": "[title]"
}
```

you can also update content in same way

### Data example

```json
{
  "title": "note123",
  "content": "content 1"
}
```

### Success Response

**Code** : `201 Created`

### Content example

```json
{
  "data": {
    "message": "Successfully updated"
  }
}
```

---

## DELETE

Used to get the a note by its ID

**URL** : `/notes/:id`

**Method** : `DELETE`

**Auth required** : NO

**Data constraints required** : NO

### Success Response

**Code** : `204 Content`
