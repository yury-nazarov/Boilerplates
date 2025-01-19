# PlantUML

1. Install PlugIn PlantUML for IDE
2. Install graphviz `brew install graphviz`
3. Check `dot -V`


## Includes

```bash
# From local file
!include ../template/C4_Styles.puml

# by link to remote dir
!define AWSPuml https://raw.githubusercontent.com/awslabs/aws-icons-for-plantuml/v18.0/dist
!include AWSPuml/AWSCommon.puml

# by link to remote file
!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml
```

## Links
- [https://c4model.com/]()
- [https://plantuml.com/ru/guide]()
- [PlantUML Shapes](https://github.com/plantuml-stdlib/C4-PlantUML?tab=readme-ov-file)
- [AWS icon for plantuml](https://github.com/awslabs/aws-icons-for-plantuml)
