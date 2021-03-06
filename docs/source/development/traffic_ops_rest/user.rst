.. 
.. Copyright 2015 Comcast Cable Communications Management, LLC
.. 
.. Licensed under the Apache License, Version 2.0 (the "License");
.. you may not use this file except in compliance with the License.
.. You may obtain a copy of the License at
.. 
..     http://www.apache.org/licenses/LICENSE-2.0
.. 
.. Unless required by applicable law or agreed to in writing, software
.. distributed under the License is distributed on an "AS IS" BASIS,
.. WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
.. See the License for the specific language governing permissions and
.. limitations under the License.
.. 

.. _to-api-users:

Users
=====

**GET /api/1.1/users.json**

Retrieves all users.

Authentication Required: Yes

Response Content Type: application/json

**Response Messages**

::


  HTTP Status Code: 200
  Reason: Success

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``email``             | string |                                                |
+----------------------+--------+------------------------------------------------+
|``city``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``id``                | hash   |                                                |
+----------------------+--------+------------------------------------------------+
|``phoneNumber``       | string |                                                |
+----------------------+--------+------------------------------------------------+
|``company``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``country``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``fullName``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``localUser``         | string |                                                |
+----------------------+--------+------------------------------------------------+
|``uid``               | string |                                                |
+----------------------+--------+------------------------------------------------+
|``username``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``rolename``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``newUser``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``addressLine2``      | string |                                                |
+----------------------+--------+------------------------------------------------+
|``role``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``addressLine1``      | string |                                                |
+----------------------+--------+------------------------------------------------+
|``postalCode``        | string |                                                |
+----------------------+--------+------------------------------------------------+
|``gid``               | string |                                                |
+----------------------+--------+------------------------------------------------+


**Response Example**

::


  [
   {
      "email": "email@email.com",
      "city": "",
      "id": "54",
      "phoneNumber": "",
      "company": "",
      "country": "",
      "fullName": "Bob Simpson",
      "localUser": false,
      "uid": "0",
      "stateOrProvince": "",
      "username": "bsimpson",
      "rolename": "portal",
      "newUser": true,
      "addressLine2": "",
      "role": "6",
      "addressLine1": "",
      "postalCode": "",
      "gid": "0"
   }
  ]

For error messages, see :ref:`reference-label-401`.

|

**GET /api/1.1/user/current.json**

Retrieves the profile for the authenticated user.

Authentication Required: Yes

Response Content Type: application/json

**Response Messages**

::


  HTTP Status Code: 200
  Reason: Success

**Request Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``email``             | string |                                                |
+----------------------+--------+------------------------------------------------+
|``city``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``id``                | string |                                                |
+----------------------+--------+------------------------------------------------+
|``phoneNumber``       | string |                                                |
+----------------------+--------+------------------------------------------------+
|``company``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``country``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``fullName``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``localUser``         | boolean|                                                |
+----------------------+--------+------------------------------------------------+
|``uid``               | string |                                                |
+----------------------+--------+------------------------------------------------+
|``stateOrProvince``   | string |                                                |
+----------------------+--------+------------------------------------------------+
|``username``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``newUser``           | boolean|                                                |
+----------------------+--------+------------------------------------------------+
|``addressLine2``      | string |                                                |
+----------------------+--------+------------------------------------------------+
|``role``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``addressLine1``      | string |                                                |
+----------------------+--------+------------------------------------------------+
|``gid``               | string |                                                |
+----------------------+--------+------------------------------------------------+
|``postalCode``        | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Example**

::

  {
         “response”: {
                          “email”: "email@email.com",
                          “city”: "",
                          “id”: "50",
                          “phoneNumber”: "",
                          “company”: "",
                          “country”: "",
                          “fullName”: "Tom Callahan",
                          “localUser”: true,
                          “uid”: "0",
                          “stateOrProvince”: "",
                          “username”: "tommyboy",
                          “newUser”: false,
                          “addressLine2”: "",
                          “role”: "6",
                          “addressLine1”: "",
                          “gid”: "0",
                          “postalCode”: ""
         },
         “version”: "1.1"
  }

For error messages, see :ref:`reference-label-401`.

|
  
**POST /api/1.1/user/current/update**

Retrieves the profile for the authenticated user.

Authentication Required: Yes

Response Content Type: application/json

**Request Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``email``             | string |                                                |
+----------------------+--------+------------------------------------------------+
|``city``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``id``                | string |                                                |
+----------------------+--------+------------------------------------------------+
|``phoneNumber``       | string |                                                |
+----------------------+--------+------------------------------------------------+
|``company``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``country``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``fullName``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``localUser``         | boolean|                                                |
+----------------------+--------+------------------------------------------------+
|``uid``               | string |                                                |
+----------------------+--------+------------------------------------------------+
|``stateOrProvince``   | string |                                                |
+----------------------+--------+------------------------------------------------+
|``username``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``newUser``           | boolean|                                                |
+----------------------+--------+------------------------------------------------+
|``addressLine2``      | string |                                                |
+----------------------+--------+------------------------------------------------+
|``role``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``addressLine1``      | string |                                                |
+----------------------+--------+------------------------------------------------+
|``gid``               | string |                                                |
+----------------------+--------+------------------------------------------------+
|``postalCode``        | string |                                                |
+----------------------+--------+------------------------------------------------+

**Request Example**


::


  {
   "user": {
      "email": "",
      "city": "",
      "id": "",
      "phoneNumber": "",
      "company": "",
      "country": "",
      "fullName": "",
      "localUser": true,
      "uid": "0",
      "stateOrProvince": "",
      "username": "tommyboy",
      "newUser": false,
      "addressLine2": "",
      "role": "6",
      "addressLine1": "",
      "gid": "0",
      "postalCode": ""
   }
  }

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``alerts``            | array  | A collection of alert messages.                |
+----------------------+--------+------------------------------------------------+
|> ``level``           | string | Success, info, warning or error.               |
+----------------------+--------+------------------------------------------------+
|> ``text``            | string | Alert message.                                 |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Messages and Examples**

::

  HTTP Status Code: 200
  Reason: Success

  {
        "alerts": [
                  {
                          "level": "success",
                          "text": "UserProfile was successfully updated."
                  }
          ],
          "version": "1.1"
  }

For error messages, see :ref:`reference-label-400` and :ref:`reference-label-401`.

**GET /api/1.1/user/current/jobs.json**

Retrieves user purge jobs.

Authentication Required: Yes

Response Content Type: application/json

**Response Messages**

::


  HTTP Status Code: 200
  Reason: Success

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``keyword``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``objectName``        | string |                                                |
+----------------------+--------+------------------------------------------------+
|``assetUrl``          | string |                                                |
+----------------------+--------+------------------------------------------------+
|``assetType``         | string |                                                |
+----------------------+--------+------------------------------------------------+
|``status``            | string |                                                |
+----------------------+--------+------------------------------------------------+
|``dsId``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``dsXmlId``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``username``          | boolean|                                                |
+----------------------+--------+------------------------------------------------+
|``parameters``        | string |                                                |
+----------------------+--------+------------------------------------------------+
|``enteredTime``       | string |                                                |
+----------------------+--------+------------------------------------------------+
|``objectType``        | string |                                                |
+----------------------+--------+------------------------------------------------+
|``agent``             | string |                                                |
+----------------------+--------+------------------------------------------------+
|``id``                | string |                                                |
+----------------------+--------+------------------------------------------------+
|``startTime``         | string |                                                |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Example**
::

  {
   "response": [
      {
         "id": "1",
         "keyword": "PURGE",
         "objectName": null,
         "assetUrl": "",
         "assetType": "file",
         "status": "PENDING",
         "dsId": "73",
         "dsXmlId": "cim-jitp",
         "username": "peewee",
         "parameters": "TTL:56h",
         "enteredTime": "2015-01-21 18:00:16",
         "objectType": null,
         "agent": "",
         "startTime": "2015-01-21 10:45:38"
      }
   ],
   "version": "1.1"
  }

For error responses, see :ref:`to-api-error`.


|

**POST/api/1.1/user/current/jobs**

Creates a purge job.

Authentication Required: Yes


**Request Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``dsId``              | string |                                                |
+----------------------+--------+------------------------------------------------+
|``dsXmlId``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|``regex``             | string |                                                |
+----------------------+--------+------------------------------------------------+
|``startTime``         | string |                                                |
+----------------------+--------+------------------------------------------------+
|``ttl``               | int    |                                                |
+----------------------+--------+------------------------------------------------+

**Request Example**

::

  {
         "dsId": "73",
         "dsXmlId": "cim-jitp",
         "regex": "/path/to/content.jpg",
         "startTime": "2015-01-27 11:08:37",
         "ttl": 54
  }

Response Content Type: application/json

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``alerts``            | array  | A collection of alert messages.                |
+----------------------+--------+------------------------------------------------+
|> ``level``           | string | Success, info, warning or error.               |
+----------------------+--------+------------------------------------------------+
|> ``text``            | string | Alert message.                                 |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Messages and Examples**

::


  HTTP Status Code: 200
  Reason: Success

  {
        “alerts”:
                [
                    { 
                          “level”: "success",
                          “text”: "Successfully created purge job for: ."
                    }
                ],
        “version”: "1.1"
  }

For error messages, see :ref:`reference-label-400` and :ref:`reference-label-401`.

|

**POST /api/1.1/user/login { u: '', p: '' }**

Authentication of a user using username and password.

Authentication Required: No

**Request Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``u``                 | string | username                                       |
+----------------------+--------+------------------------------------------------+
|``p``                 | string | password                                       |
+----------------------+--------+------------------------------------------------+

**Request Example**

::

  {
   "u": "username",
   "p": "password"
 }

Response Content Type: application/json

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``alerts``            | array  | A collection of alert messages.                |
+----------------------+--------+------------------------------------------------+
|> ``level``           | string | Success, info, warning or error.               |
+----------------------+--------+------------------------------------------------+
|> ``text``            | string | Alert message.                                 |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Messages and Examples**

::


  HTTP Status Code: 200
  Reason: Success

 {
   "alerts": [
      {
         "level": "success",
         "text": "Successfully logged in."
      }
   ],
   "version": "1.1"
  }

For error messages, see :ref:`reference-label-401`.

|

**GET /api/1.1/user/:id/deliveryservices/available.json**

Authentication Required: Yes

Response Content Type: application/json

**Request Route Parameters**

+-----------------+----------+---------------------------------------------------+
| Name            | Required | Description                                       |
+=================+==========+===================================================+
|id               | yes      |                                                   |
+-----------------+----------+---------------------------------------------------+

**Response Messages**

::


  HTTP Status Code: 200
  Reason: Success

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``xmlId``             | string |                                                |
+----------------------+--------+------------------------------------------------+
|``id``                | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Example**


::


  {
   "response": [
      {
         "xmlId": "ns-img",
         "id": "90"
      },
      {
         "xmlId": "ns-img-secure",
         "id": "280"
      }
   ],
   "version": "1.1"
  }

For error messages, see :ref:`reference-label-401`.

|

**POST /api/1.1/user/login/token**

Authentication of a user using a temporary token.

Authentication Required: No

**Request Route Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``t``                 | string | token-value                                    |
+----------------------+--------+------------------------------------------------+

**Request Example**

::


  {
   "t": "token-value"
  }

Response Content Type: application/json

**Response Messages**

::


  HTTP Status Code: 200
  Reason: Success

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``alerts``            | array  |                                                |
+----------------------+--------+------------------------------------------------+
|> ``level``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|> ``text``            | string |                                                |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Example**

::

  {
   "alerts": [
      {
         "level": "error",
         "text": "Unauthorized, please log in."
      }
   ],
   "version": "1.1"
  }

For error messages, see :ref:`reference-label-401`.

|


**POST /api/1.1/user/logout**

User logout.

Authentication Required: Yes

Response Content Type: application/json

**Response Messages**

::


  HTTP Status Code: 200
  Reason: Success

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``alerts``            | array  |                                                |
+----------------------+--------+------------------------------------------------+
|* ``level``           | string |                                                |
+----------------------+--------+------------------------------------------------+
|* ``text``            | string |                                                |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Example**

::

  {
   "alerts": [
      {
         "level": "success",
         "text": "You are logged out."
      }
   ],
   "version": "1.1"
  }

For error messages, see :ref:`reference-label-401`.

|

**POST /api/1.1/user/reset_password**

Reset user password.

Authentication Required: No

**Request Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``email``             | string | The email address of the user to initiate      |
|                      |        | password reset.                                |
+----------------------+--------+------------------------------------------------+

**Request Example**

::

  {
   "email": "email@email.com"
  }

Response Content Type: application/json

**Response Properties**

+----------------------+--------+------------------------------------------------+
| Parameter            | Type   | Description                                    |
+======================+========+================================================+
|``alerts``            | array  | A collection of alert messages.                |
+----------------------+--------+------------------------------------------------+
|* ``level``           | string | Success, info, warning or error.               |
+----------------------+--------+------------------------------------------------+
|* ``text``            | string | Alert message.                                 |
+----------------------+--------+------------------------------------------------+
|``version``           | string |                                                |
+----------------------+--------+------------------------------------------------+

**Response Messages and Examples**

::

  
  HTTP Status Code: 200
  Reason: Success

  {
   "alerts": [
      {
         "level": "success",
         "text": "Successfully logged in."
      }
   ],
   "version": "1.1"
  }

For error messages, see :ref:`reference-label-400`   
  