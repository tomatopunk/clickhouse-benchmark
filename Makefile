LICENSE_HEADER := "/*" \
	"\n * Licensed to the Apache Software Foundation (ASF) under one" \
	"\n * or more contributor license agreements.  See the NOTICE file" \
	"\n * distributed with this work for additional information" \
	"\n * regarding copyright ownership.  The ASF licenses this file" \
	"\n * to you under the Apache License, Version 2.0 (the" \
	"\n * \"License\"); you may not use this file except in compliance" \
	"\n * with the License.  You may obtain a copy of the License at" \
	"\n *" \
	"\n *   http://www.apache.org/licenses/LICENSE-2.0" \
	"\n *" \
	"\n * Unless required by applicable law or agreed to in writing," \
	"\n * software distributed under the License is distributed on an" \
	"\n * \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY" \
	"\n * KIND, either express or implied.  See the License for the" \
	"\n * specific language governing permissions and limitations" \
	"\n * under the License." \
	"\n */"

GO_FILES := $(shell find . -type f -name "*.go")

.PHONY: add-license

add-license:
	@for file in $(GO_FILES); do \
		echo "Adding license to $$file"; \
		tmpfile=$$(mktemp); \
		echo -e $(LICENSE_HEADER) | cat - $$file > $$tmpfile && mv $$tmpfile $$file; \
	done
