
# CoMID Template Format

## 1 Introduction

**CoMID (Concise Model Identifier)**,  is a structured tag that encapsulates detailed information about the composition of hardware, firmware, or modules. Each CoMID is uniquely identified by a specific ID, enabling unambiguous reference to CoMID instances, particularly in contexts such as typed link relations or CoBOM (Concise Bill of Material) tags.

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

![CoMID](../../data/pics/CoMID.png)

## 6. Full Examples and Walkthroughs

Below are the **seven** template files, each highlighting different aspects of CoMID usage.

### 6.1. comid-cca-mult-refval.json
```
{
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B16",
    "version": 0
  },
  "entities": [
    {
      "name": "ACME Ltd.",
      "regid": "https://acme.example",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          }
        },
        "measurements": [
          {
            "key": {
              "type": "psa.refval-id",
              "value": {
                "label": "BL",
                "version": "2.1.0",
                "signer-id": "rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="
              }
            },
            "value": {
              "digests": [
                "sha-256:h0KPxSKAPTEGXnvOPPA/5HUJZjHl4Hu9eg/eYMTPJcc="
              ]
            }
          },
          {
            "key": {
              "type": "psa.refval-id",
              "value": {
                "label": "PRoT",
                "version": "1.3.5",
                "signer-id": "rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="
              }
            },
            "value": {
              "digests": [
                "sha-256:AmOCmYm2/ZVPcrqvL8ZLwuLwHWktTecphuqAj26ZgT8="
              ]
            }
          },
          {
            "key": {
              "type": "psa.refval-id",
              "value": {
                "label": "ARoT",
                "version": "0.1.4",
                "signer-id": "rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="
              }
            },
            "value": {
              "digests": [
                "sha-256:o6XnFfDMV0pzw/m+u2vCTzL/1bZ7OHJEwskJ2neaFHg="
              ]
            }
          },
          {
            "key": {
              "type": "cca.platform-config-id",
              "value": "cfg v1.0.0"
            },
            "value": {
              "raw-value": {
                "type": "bytes",
                "value": "cmF3dmFsdWUKcmF3dmFsdWUK"
              }
            }
          }
        ]
      }
    ]
  }
}
``` 

</details>

**Key Points**

-   **Multiple Reference Values**: Demonstrates multiple reference values within a single environment.
-   **Identifier Types**: Uses `psa.impl-id` for environment identifier.
-   **Diverse Measurement Keys**: Showcases different measurement keys (`psa.refval-id` and `cca.platform-config-id`).
-   **Raw Values**: Includes raw byte values for platform configuration.

### 6.2. comid-cca-realm-refval.json
```
{
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B16",
    "version": 0
  },
  "entities": [
    {
      "name": "Workload Client Ltd.",
      "regid": "https://workloadclient.example",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "uuid",
              "value": "CD1F0E55-26F9-460D-B9D8-F7FDE171787C"
            },
            "vendor": "Workload Client Ltd"
          },
          "instance": {
            "type": "bytes",
            "value": "QoS1aUymwNLPR4mguVrIAlyBjeUjBDZL580pgbLS7caFsyInfsJYGZYkE9jJssH1"
          }
        },
        "measurements": [
          {
            "value": {
              "integrity-registers": {
                "rim": {
                  "key-type": "text",
                  "value": [
                    "sha-384;QoS1aUymwNLPR4mguVrIAlyBjeUjBDZL580pgbLS7caFsyInfsJYGZYkE9jJssH1"
                  ]
                },
                "rem0": {
                  "key-type": "text",
                  "value": [
                    "sha-384;IQe752H8pS2VE2oTVNt6TdV7Gya+DT2nHZ6yOYazS6YVq/ZRTPNeWp6lWgMtBop4"
                  ]
                },
                "rem1": {
                  "key-type": "text",
                  "value": [
                    "sha-384;JQe752H8pS2VE2oTVNt6TdV7Gya+DT2nHZ6yOYazS6YVq/ZRTPNeWp6lWgMtBop4"
                  ]
                },
                "rem2": {
                  "key-type": "text",
                  "value": [
                    "sha-384;MQe752H8pS2VE2oTVNt6TdV7Gya+DT2nHZ6yOYazS6YVq/ZRTPNeWp6lWgMtBop4"
                  ]
                },
                "rem3": {
                  "key-type": "text",
                  "value": [
                    "sha-384;NQe752H8pS2VE2oTVNt6TdV7Gya+DT2nHZ6yOYazS6YVq/ZRTPNeWp6lWgMtBop4"
                  ]
                }
              }
            }
          }
        ]
      }
    ]
  }
}
``` 
**Key Points**

-   **UUID Identifiers**: Utilizes `uuid` type for environment identifiers, ensuring uniqueness.
-   **Instance Field**: Includes an `instance` field of type `bytes` to differentiate instances.
-   **Integrity Registers**: Uses `integrity-registers` to hold multiple cryptographic hash values across different measurement registers (`rim`, `rem0`, `rem1`, etc.).
-   **Operational Flags**: Contains `op-flags` indicating properties like `notSecure` and `debug`.
-   **Software Version Number (svn)**: Specifies software versioning with exact values.

### 6.3 comid-cca-refval.json
``` 
{
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B16",
    "version": 0
  },
  "entities": [
    {
      "name": "ACME Ltd.",
      "regid": "https://acme.example",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          }
        },
        "measurements": [
          {
            "key": {
              "type": "cca.platform-config-id",
              "value": "cfg v1.0.0"
            },
            "value": {
              "raw-value": {
                "type": "bytes",
                "value": "cmF3dmFsdWUKcmF3dmFsdWUK"
              }
            }
          }
        ]
      }
    ]
  }
}
``` 

**Key Points**

-   **Multiple Reference-Value Entries**: Illustrates handling of multiple reference values within a single CoMID.
-   **Operational Flags**: Uses `op-flags` to indicate operational states like `notSecure` and `debug`.
-   **Software Version Numbers (svn)**: Uses `svn` to specify exact software versions for each measurement.
-   **Layered Environments**: Demonstrates multi-layered environments (e.g., layer 0 and layer 1) with unique identifiers.
-   **Digest Arrays**: Shows multiple digest entries within a single measurement for enhanced integrity verification.

### 6.4 comid-dice-refval.json

```
{
  "tag-identity": {
    "id": "1D5A8C7C-1C70-4C56-937E-3C5713AE5A83"
  },
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "uuid",
              "value": "DD6661F0-0928-4401-966B-589EA74E3272"
            },
            "model": "FMC",
            "layer": 0,
            "index": 0
          }
        },
        "measurements": [
          {
            "value": {
              "op-flags": [
                "notSecure",
                "debug"
              ],
              "digests": [
                "sha-256:RKozavTLFKh5Qy5T3WVxx/qbzK+3X0iCWSYtbqOk2Rs="
              ],
              "svn": {
                "type": "exact-value",
                "value": 10
              }
            }
          }
        ]
      },
      {
        "environment": {
          "class": {
            "id": {
              "type": "uuid",
              "value": "FFDA7CF3-2333-4A91-99A8-068626203ACA"
            },
            "model": "L1",
            "layer": 1,
            "index": 0
          }
        },
        "measurements": [
          {
            "value": {
              "op-flags": [
                "notSecure",
                "debug"
              ],
              "digests": [
                "sha-256:h0KPxSKAPTEGXnvOPPA/5HUJZjHl4Hu9eg/eYMTPJcc=",
                "sha-256:VgXOanU71cskR7hhl418y0an8zsD772wLJYg2o6awD0="
              ],
              "svn": {
                "type": "exact-value",
                "value": 2
              }
            }
          }
        ]
      }
    ]
  }
}
``` 

**Key Points**

-   **DICE Framework Integration**: The file name `comid-dice-refval.json` suggests integration with the DICE framework, which utilizes layered measurements for enhanced security.
-   **Layered Measurements**: Contains multiple layers (`layer`: 0 and `layer`: 1), each with distinct environments and measurements.
-   **Operational Flags**: Includes flags like `notSecure` and `debug` to indicate the operational state of each environment.
-   **Software Version Numbers (svn)**: Uses `svn` to specify exact software versions for each measurement.
-   **Multiple Digests**: Some measurements contain multiple digests to strengthen integrity verification.

### 6.5 comid-psa-iakpub.json

``` 
{
  "lang": "en-GB",
  "tag-identity": {
    "id": "366D0A0A-5988-45ED-8488-2F2A544F6242",
    "version": 0
  },
  "entities": [
    {
      "name": "ACME Ltd.",
      "regid": "https://acme.example",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "attester-verification-keys": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          },
          "instance": {
            "type": "ueid",
            "value": "Ac7rrnuJJ6MiflMDz14PH3s0u1Qq1yUKwD+83jbsLxUI"
          }
        },
        "verification-keys": [
          {
            "type": "pkix-base64-key",
            "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFn0taoAwR3PmrKkYLtAsD9o05KSM6mbgfNCgpuL0g6VpTHkZl73wk5BDxoV7n+Oeee0iIqkW3HMZT3ETiniJdg==\n-----END PUBLIC KEY-----"
          }
        ]
      },
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          },
          "instance": {
            "type": "ueid",
            "value": "AUyj5PUL8kjDl4cCDWj/0FyIdndRvyZFypI/V6mL7NKW"
          }
        },
        "verification-keys": [
          {
            "type": "pkix-base64-key",
            "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6Vwqe7hy3O8Ypa+BUETLUjBNU3rEXVUyt9XHR7HJWLG7XTKQd9i1kVRXeBPDLFnfYru1/euxRnJM7H9UoFDLdA==\n-----END PUBLIC KEY-----"
          }
        ]
      }
    ]
  }
}
``` 
**Key Points**

-   **Public Keys for Attestation**: Stores two **verification-keys** under different **instances** (`ueid`: Unique Entity Identifier), enabling attesters to verify measurements or claims.
-   **Key Formats**: Utilizes `pkix-base64-key` type, ensuring keys are in a standard PEM format.
-   **Consistent Environment Class**: Both verification keys are associated with the same `class` (`psa.impl-id`, `vendor`, `model`), indicating they belong to the same hardware or firmware class.
-   **Unique Instances**: Differentiates verification keys using unique `ueid` values, allowing multiple keys per environment.

### 6.6 comid-psa-integ-iakpub.json
```
{
  "lang": "en-GB",
  "tag-identity": {
    "id": "366D0A0A-5988-45ED-8488-2F2A544F6242",
    "version": 0
  },
  "entities": [
    {
      "name": "ACME Ltd.",
      "regid": "https://acme.example",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "attester-verification-keys": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          },
          "instance": {
            "type": "ueid",
            "value": "Ac7rrnuJJ6MiflMDz14PH3s0u1Qq1yUKwD+83jbsLxUI"
          }
        },
        "verification-keys": [
          {
            "type": "pkix-base64-key",
            "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMKBCTNIcKUSDii11ySs3526iDZ8A\niTo7Tu6KPAqv7D7gS2XpJFbZiItSs3m9+9Ue6GnvHw/GW2ZZaVtszggXIw==\n-----END PUBLIC KEY-----"
          }
        ]
      },
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          },
          "instance": {
            "type": "ueid",
            "value": "AUyj5PUL8kjDl4cCDWj/0FyIdndRvyZFypI/V6mL7NKW"
          }
        },
        "verification-keys": [
          {
            "type": "pkix-base64-key",
            "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6Vwqe7hy3O8Ypa+BUETLUjBNU3rEXVUyt9XHR7HJWLG7XTKQd9i1kVRXeBPDLFnfYru1/euxRnJM7H9UoFDLdA==\n-----END PUBLIC KEY-----"
          }
        ]
      }
    ]
  }
}
``` 
**Key Points**

-   **Integration with PSA**: The file name indicates integration with the Platform Security Architecture (PSA), emphasizing secure key management.
-   **Multiple Verification Keys**: Similar to `comid-psa-iakpub.json`, this file includes multiple verification keys, each associated with distinct `ueid` instances.
-   **Key Material Variations**: Demonstrates how different environments can have unique key materials, enhancing security through key separation.
-   **Consistent Class Definition**: Maintains the same `class` details across verification keys, ensuring they are recognized as part of the same environment.

### 6.7 comid-psa-refval.json
```
{
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B16",
    "version": 0
  },
  "entities": [
    {
      "name": "ACME Ltd.",
      "regid": "https://acme.example",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "psa.impl-id",
              "value": "YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="
            },
            "vendor": "ACME",
            "model": "RoadRunner"
          }
        },
        "measurements": [
          {
            "key": {
              "type": "psa.refval-id",
              "value": {
                "label": "BL",
                "version": "2.1.0",
                "signer-id": "rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="
              }
            },
            "value": {
              "digests": [
                "sha-256:h0KPxSKAPTEGXnvOPPA/5HUJZjHl4Hu9eg/eYMTPJcc="
              ]
            }
          },
          {
            "key": {
              "type": "psa.refval-id",
              "value": {
                "label": "PRoT",
                "version": "1.3.5",
                "signer-id": "rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="
              }
            },
            "value": {
              "digests": [
                "sha-256:AmOCmYm2/ZVPcrqvL8ZLwuLwHWktTecphuqAj26ZgT8="
              ]
            }
          },
          {
            "key": {
              "type": "psa.refval-id",
              "value": {
                "label": "ARoT",
                "version": "0.1.4",
                "signer-id": "rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="
              }
            },
            "value": {
              "digests": [
                "sha-256:o6XnFfDMV0pzw/m+u2vCTzL/1bZ7OHJEwskJ2neaFHg="
              ]
            }
          }
        ]
      }
    ]
  }
}
```

**Key Points**

-   **Consistent Reference Values**: All reference values utilize the `psa.refval-id` type, maintaining consistency in measurement identification.
-   **Uniform Digest Algorithm**: Employs `sha-256` across all measurements, ensuring uniform integrity verification.
-   **Signer Identification**: Each measurement includes a `signer-id`, linking the digest to its trusted signer.
-   **Standard Structure**: Adheres to the standard CoMID structure for PSA-based reference values, facilitating interoperability and ease of verification.


### 7 Lifecycle of a CoMID Templates
![CoMID_Lifecycle](../../data/pics/comid life.png)
