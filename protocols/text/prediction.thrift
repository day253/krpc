namespace go shumei.strategy.re

struct Text {
    1: optional binary content,
}

struct TextPredictRequest {
    1: optional string requestId,  
    2: optional string organization,  
    3: optional string product,
    4: optional string sceneId,
    5: optional string appId,
    6: optional string channel,
    7: optional string data,
    8: optional Text text,  
}

struct TextPredictResult {
    1: optional string result,
}

service TextPredictor {
    TextPredictResult predict(1: TextPredictRequest request),
    bool health(),
}
