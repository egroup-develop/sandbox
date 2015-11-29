#-*- coding: utf-8 -*-
#本スクリプトの実行時に第一引数として特徴を抽出したい画像の相対パスを与える

import caffe
from caffe.proto import caffe_pb2
import numpy as np
import sys

#ここではip2層の値を取り出している. outputをクラス数分の100にしたので100個特徴量が得られる
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
#与えられた画像が属するクラス(最も近いクラス)を表示
pred = np.argmax(predictions)

#caffeで特徴抽出から分類まで完結
print predictions
print "\n" + sys.argv[1] + "が属するクラス(一番近い人)は, "
print pred
print "です"

#内積層の最後の層の出力を表示(num_output分)する
#index: 素性(特徴量) の組
feat = classifier.blobs[LAYER].data[INDEX].flatten().tolist()

#近しいもの
featSorted = feat
featSorted = np.sort(featSorted)
featSorted = featSorted[-1::-1]
pred2 = featSorted[1]
pred3 = featSorted[2]

for index, value in enumerate(feat):
  if pred2 == value:
    print "2番目に近い人は, "
    print index
    print "です"
  elif pred3 == value:
    print "3番目に近い人は, "
    print index
    print "です"
print "\n"

print sys.argv[1] + "の特徴量..."

#for i,f in enumerate(feat): 
#  print(str(i+1) + ":" + str(f)),
for i,f in enumerate(feat):
  print(str(i+1) + ":" + str(f) + "\n"),
