
Default values are used when provided at the top of the schema:

```yaml
  schema:
    type: object
    default:
      foo: true
      bar: 42
      bam: boom
    properties:
      foo: 
        type: boolean
      bar:
        type: integer
      bam:
        type: string
```

Code generated with the the 0.13.0 release.

    make clean generate
