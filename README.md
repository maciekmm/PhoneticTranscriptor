#Phonetic transcriptor
####Transcripts phonetic text to language specific phones!
*trænskrɪpts fənɛtɪk tɛkst tu læŋgwədʒ spəsɪfɪk fonz!*<br />*transcripts feaneatic teacst tu langweaj speasific fonz!*

----
## Requirements
1. GO package

## Usage
1. Get phonetic transcriptor
```go get github.com/maciekmm/phonetictranscriptor```
2. Copy languages files to your executable directory (will be replaced by packing mechanism later)
3. Import and use it!
```
import "phonetictranscriptor"
func sth() {
    phonetictranscriptor.TranscriptByCode("pl","tajp jɔr tɛkst hɪr.")
}
```
