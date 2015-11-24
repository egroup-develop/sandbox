#-*- coding: utf-8 -*-
#本スクリプトの実行時に第一引数として特徴を抽出したい画像の相対パスを与える

import caffe
from caffe.proto import caffe_pb2
import numpy as np
import sys

#ここではip2層の値を取り出している
LAYER = 'ip2'
INDEX = 0

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

#iamge...image's path(相対でよい)
#image = caffe.io.load_image('./LogirlImages/0/image_1.png')
image = caffe.io.load_image(sys.argv[1])
#特徴量
predictions = classifier.predict([image], oversample=False)
#与えられた画像が属するクラスを表示
pred = np.argmax(predictions)

#caffeで特徴抽出から分類まで完結
print predictions
print sys.argv[1] + "が属するクラスは, \n"
print pred
print "です\n"

print sys.argv[1] + "の特徴量..."

#caffeを特徴抽出として使う用
#index: 素性(特徴量) の組を出力
feat = classifier.blobs[LAYER].data[INDEX].flatten().tolist()
for i,f in enumerate(feat): 
  print(str(i+1) + ":" + str(f)),
