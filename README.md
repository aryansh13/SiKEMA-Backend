
## API Reference

#### Get all event

```http
  GET /api/lecturer/{lecturerid}/event
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Query  | Description                |
| :-------- | :------------------------- |
| `page` | Pagination page  |
| `itemperpage` | Controls how many item displayed per page |

#### Get event

```http
  GET /api/lecturer/{lecturerid}/event/${id}
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Query  | Description                |
| :-------- | :------------------------- |
| `page` | Pagination page  |
| `itemperpage` | Controls how many item displayed per page |

#### Create new event

```http
  POST /api/lecturer/{lecturerid}/event
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `course_id` | `uint` | **Required**. Course ID |
| `class_id` | `uint` | **Required**. Class ID intended for attendance  |
| `meet` | `uint` | **Required**. Indicates which course meetings |

#### Record attendance (add student to attendance event)

```http
  POST /api/lecturer/{lecturerid}/event/{id}/students
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `student_id`      | `array[string]` | **Required**. List of IDs of the student whose attendance will be recorded |

#### Remove attendance (remove student from attendance event)

```http
  DELETE /api/lecturer/{lecturerid}/event/{id}/students
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `student_id`      | `array[string]` | **Required**. List of IDs of the student whose attendance will be removed |

#### Finalize event (Event will now be read-only)

```http
  PATCH/PUT /api/lecturer/{lecturerid}/event/{id}
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `status`      | `uint` | **Required**. Attendance event status |

