#!/usr/bin/env python 
# -*- python -*-
#BEGIN_LEGAL
#
#Copyright (c) 2024 Intel Corporation
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.
#  
#END_LEGAL
from __future__ import print_function
import sys
import os
from stat import *

def find_dir(d):
    dir = os.getcwd()
    last = ''
    while dir != last:
        target_dir = os.path.join(dir,d)
        if os.path.exists(target_dir):
            return target_dir
        last = dir
        (dir,tail) = os.path.split(dir)
    return None

XED_ROOT = find_dir('xed')
sys.path.append(XED_ROOT)

from scripts import apply_legal_header 

apply_legal_header.HEADER_SKIP_COMMITS = ['36e5f9a51b59c2a1ebeca9dcf674e0817b514a5a']

if __name__ == '__main__':

    apply_legal_header.main()
