namespace go shumei.strategy.re

struct Event {
}

struct EventPredictRequest {
    1: optional string requestId, 
    2: optional string organization, 
    3: optional Event event,
}

struct EventPredictResult {
    1: optional string result,  
}

service EventPredictor {
    EventPredictResult predict(1: EventPredictRequest request),
    bool health(),
}
