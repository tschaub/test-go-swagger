
The required property works only if there is no default object.

```yaml
  schema:
    type: object
    required:
      - chicken
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

    make clean generate
