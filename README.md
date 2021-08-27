# admiral-template

A template on how we think admiral can be used best.

## Directory structure

```shell
├── cmd         # your binaries
│ └── cli       # entrypoint
├── internal    # your app business case
│ ├── actions   # command actions
│ ├── commands  # the app commands
│ ├── hooks     # your command hooks (pre & post run)
│ ├── ns        # Namespaces for your commands.
│ └── services  # services for your commands
└── pkg         # reusable code
└── utils       # global utilities for your implementation
```

## Modifiying the binary output.
