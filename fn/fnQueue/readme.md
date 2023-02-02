# MiniQueue


## 쓰기작업 큐 정리
~~~mermaid
stateDiagram-v2
direction LR

[*] --> Request1
[*] --> Request2
[*] --> Request3

Request2
Request3

state Request1 {
    RequestRecord1
    UsingQueueNumber1
    
    note right of UsingQueueNumber1
        미리 로드벨런싱된 큐 번호가 정해져 있다.
        즉, 한개의 데이터는 이미 지정된 큐만 쓰게 되므로
        다중 서비스(스케일링) 되더래도, 한개의 데이터는
        반드시 순서대로 처리되게 된다.
        결과적으로 데이터는 자기 자신이 구독할 큐의 정보를 가지고 있어야 한다.
        이후 큐 사용량에 대한 분석을 통해 리벨런싱해서 전체 적인 시스템의 분산을 해주는 보조 프로그램까지 개발하면
        분산처리도 완료되게 된다.
    end note
}

Request1 --> Func1_1
Func1_1 --> Queue1
Queue1 --> Func1_2
Func1_2 --> Record1 : Request

state Service {
    state FrontService {
        Func1_1
        Func2_1
        Func3_1
    }
    
    state BackService {
        Func1_2
        Func2_2
        Func3_2
    }
}
state RequestQueue {
    Queue1
}
  
state Database {
    state Record1 {
        Data
        UsingQueueNumber
    }
    
    Record2
    Record3
}

Record1 --> Func1_2 : Result


state ResponseQueue {
    RespQueue1
}

Func1_2 --> RespQueue1 : Publish
RespQueue1 --> Func1_1 
~~~

