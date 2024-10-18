namespace go shumei.strategy.re

struct Text {
    1: optional binary content,
}

struct TextPredictRequest {
    1: optional string requestId,  
    2: optional string organization,  
    3: optional Text text,  
}

struct TextPredictResult {
    1: optional string result,
}

service TextPredictor {
    TextPredictResult predict(1: TextPredictRequest request),
    bool health(),
}
