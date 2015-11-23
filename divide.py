# -*- coding: utf-8 -*-

import caffe
from caffe.proto import caffe_pb2
import numpy as np

mean_blob = caffe_pb2.BlobProto()
with open('mean.binaryproto') as f:
    mean_blob.ParseFromString(f.read())
mean_array = np.asarray(
mean_blob.data,
dtype=np.float32).reshape(
    (mean_blob.channels,
    mean_blob.height,
    mean_blob.width))
classifier = caffe.Classifier(
    'cifar10_quick.prototxt',
    'cifar10_quick_iter_4000.caffemodel',
    mean=mean_array,
    raw_scale=255)

#iamge...image's path
image = caffe.io.load_image('./LogirlImages/0/image_1.png')
predictions = classifier.predict([image], oversample=False)
pred = np.argmax(predictions)
print(predictions)
print(pred)
