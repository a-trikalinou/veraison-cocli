
# CoMID Template Format

## 1 Introduction

**CoMID (Concise Model Identifier)**,  is a data model and serialization format (in JSON) for capturing **reference values** and **verification keys** that can be used in remote attestation and other trust-verification scenarios. By standardizing how measurements are captured and shared, CoMID facilitates **interoperability**, **integrity**, and **traceability** across various systems and vendors.

## 2 Template Structure

A CoMID template is a JSON document composed of **top-level fields** and a **triples** object. The **top-level fields** provide overall identification, language, and authorship, while the **triples** object contains domain-specific data (e.g., reference values, attester keys).

```
{
  "lang": "<language-region>",
  "tag-identity": { ... },
  "entities": [ ... ],
  "triples": {
    "reference-values": [ ... ],
    "attester-verification-keys": [ ... ]
    ...
  }
}
``` 

### 2.1 Top-Level Fields

-   **lang** (`String`): Defines the language or locale (e.g., `"en-GB"`).
-   **tag-identity** (`Object`): Uniquely identifies this CoMID document via an ID (often a UUID) and includes a version number.
-   **entities** (`Array`): Lists the entities (organizations, individuals, etc.) contributing to or maintaining the document, along with their roles.

### 2.2 Triples

-   **reference-values**: One or more **reference-value** objects, each containing an **environment** and one or more **measurements**.
-   **attester-verification-keys**: One or more **attester-verification-key** objects, each containing an **environment** and an array of **verification-keys**.


## 3 Components

### 3.1 Environment

An **environment** captures the context of a measurement or verification key:

-   **class**: Vendor, model, and possibly an ID (`type` + `value`).
-   **instance** (`optional`): For distinguishing multiple instances of the same environment (e.g., using `ueid` or `uuid`).
-   **layer** and **index** (`optional`): For layered environments (e.g., DICE layers in multi-stage boot processes).

### 3.2 Measurements

Each measurement has two crucial subfields:

-   **key**: Identifies the measurement, including possible fields like `label`, `version`, and `signer-id`.
-   **value**: Holds the actual measurement data (e.g., cryptographic digests, raw values, or operational flags).

### 3.3 Attester Verification Keys

Used to store **public keys** associated with an environment. This is essential for verifying the attestation claims or measurement signatures produced by that environment.

## 4 Field-By-Field Explanation

### 4.1 Global Fields
|     Field    	|  Type  	|                       Description                       	|                    Example                   	|   	
|:------------:	|:------:	|:-------------------------------------------------------:	|:--------------------------------------------:	|
| lang         	| String 	| Language/country code.                                  	| "en-GB"                                      	|   	
| tag-identity 	| Object 	| Identity of this CoMID tag (UUID + version).            	| "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B16" 	|   	
| entities     	| Array  	| The organizations/roles associated with this CoMID tag. 	| [ { "name": "ACME Ltd." ... } ]              	|   	

### 4.2 Reference-Value Fields
|        Field       |  Type  |                                     Description                                    |                                            Example                                            |
|:------------------:|:------:|:----------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------:|
| environment        | Object | Contains class and optionally instance, layer, index.                              | See 3.1 Environment.                                                                          |
| measurements       | Array  | List of measurement objects.                                                       | [ { "key": { ... }, "value": { ... } } ]                                                      |
| measurements.key   | Object | Identifies the measurement. Could be a psa.refval-id, cca.platform-config-id, etc. | { "type": "psa.refval-id", "value": { "label": "BL", "version": "2.1.0", ... } }              |
| measurements.value | Object | Holds the actual measurement data.                                                 | { "digests": ["sha-256:..."] }, or { "raw-value": { "type": "bytes", "value": "..." } }, etc. |

### 4.3 Attester-Verification-Key Fields
|       Field       |  Type  |               Description               |                                   Example                                   |
|:-----------------:|:------:|:---------------------------------------:|:---------------------------------------------------------------------------:|
| environment       | Object | Defines the environment for these keys. | See 3.1 Environment.                                                        |
| verification-keys | Array  | Holds one or more public keys.          | [ { "type": "pkix-base64-key", "value": "-----BEGIN PUBLIC KEY-----..." } ] |
----------

### 5 High Level Structure for CoMID Templates

