package com.dappervision.wearscript_tagalong.events;

import android.content.Intent;

public class ActivityResultEvent {
    private int requestCode;
    private int resultCode;
    private Intent intent;


    public ActivityResultEvent(int requestCode, int resultCode, Intent intent) {
        this.requestCode = requestCode;
        this.resultCode = resultCode;
        this.intent = intent;
    }

    public int getRequestCode() {
        return requestCode;
    }

    public int getResultCode() {
        return resultCode;
    }

    public Intent getIntent() {
        return intent;
    }
}