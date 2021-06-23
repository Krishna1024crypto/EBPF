package main
import "github.com/iovisor/gobpf/bcc"

const EPFProgram='
#include<uapi/linux/ptrace.h>
BPF_PERF_OUTPUT(EVENTS)
int_handler_was called(struct reg*ctx){
char message[]="connection was estabilished";
events perf_submit(ctx,&message,sizeof(message);
return 0;
}
'
func main()
{
bpfModule=bcc.Newmodule(eBPFProgram,[]string())
uprobefd,err=bpfModule.LoadUprobe("connection was estabilished")
if err!=nil{
log.Fatal(err)
bpfModule.attachUprobes(os.Args[1],"main.handlerfunction",uprobeFd,-1)
if err!= nil{
log.Fatal(err)
table := bcc.NewTable(bpfModule.TableId("events"),bpfModule)
PerfChannel := make(chan []byte)
if err != nil{
log.Fatal(err)
}
perfMap,err=bcc.InitPerfMap(table,perfChannel,nil)
if err != nil{
log.Fatal(err)
}
perfMap.start()
deferPerfMap.stop
go fun()
for
{
values := <-perfChannel
fmt.Println(string(value))                                     
 }
 ()
 c := make(chan os.signaL,1)
 signal.Notify(c,os.Interrupt)
 <-c 
 }
