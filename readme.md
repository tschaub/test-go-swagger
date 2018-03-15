
With a schema like the one below, the default values don't appear in the generated code (except within the embedded spec).

```yaml
  schema:
    type: object
    properties:
      foo: 
        type: boolean
        default: true
      bar:
        type: integer
        default: 42 
      bam:
        type: string
        default: boom
```

Code generated with the the 0.13.0 release.

    make clean generate
