package com.dappervision.wearscript.events;

import com.thalmic.android.myo.math.Quaternion;
import com.thalmic.android.myo.math.Vector3;

public class MyoGyroDataEvent {

    private final long timestamp;
    private final Vector3 rotationRate;

    public MyoGyroDataEvent(long timestamp, Vector3 rotationRate) {
        this.timestamp = timestamp;
        this.rotationRate = rotationRate;
    }

    public long getTimestamp() {
        return timestamp;
    }

    public Vector3 getRotationRate() {
        return rotationRate;
    }
}
