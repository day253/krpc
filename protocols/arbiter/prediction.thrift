namespace go com.shumei.service

struct PredictRequest {
    1: optional string requestId;
    2: optional string organization;
    3: optional string data;
}

struct PredictResult {
    1: optional string detail;
}

service Predictor {
    PredictResult predict(1:PredictRequest request)
    bool health(),
}
