namespace go shumei.strategy.re

struct Audio {
    1: optional binary content,   // 原始二进制
    2: optional i32 sampleDepth,  // 采样深度 8bit 16bit 24bit
    3: optional i32 channelCount, // 声道数量 1 2 4 5.1 7.1 7.1.4 11.1
    4: optional i64 sampleRate,   // 采样率 44100 48000 96000 192000
    5: optional i64 duration,     // 音频时长(ms) hh::mm::ss.hhh ms
}

// 预测请求
struct AudioPredictRequest {
    1: optional string requestId,   
    2: optional string organization, 
    3: optional Audio audio,
}

struct AudioPredictResult {
    1: optional string result, 
}

service AudioPredictor {
    AudioPredictResult predict(1: AudioPredictRequest request),
    bool health(),
}
