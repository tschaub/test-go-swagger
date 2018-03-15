
Adding a list of required properties breaks if default is present as well.

```yaml
  schema:
    type: object
    required:
      - chicken
    default:
      foo: true
      bar: 42
      bam: boom
    properties:
      chicken:
        type: number
      foo: 
        type: boolean
      bar:
        type: integer
      bam:
        type: string
```

Code generated with the the 0.13.0 release.

    make validate

This results in:

    The swagger spec at "./api/spec.yml" is invalid against swagger specification 2.0. see errors :
    - body.chicken in body is required
