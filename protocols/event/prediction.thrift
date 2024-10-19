namespace go shumei.strategy.re

struct Event {
}

struct EventPredictRequest {
    1: optional string requestId, 
    2: optional string organization, 
    3: optional string product,
    4: optional string sceneId,
    5: optional string appId,
    6: optional string channel,
    7: optional string data,
    8: optional Event event,
}

struct EventPredictResult {
    1: optional string result,  
}

service EventPredictor {
    EventPredictResult predict(1: EventPredictRequest request),
    bool health(),
}
