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