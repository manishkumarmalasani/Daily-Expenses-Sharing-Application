
/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

export class Expense {
    constructor(
        public id: string,
        public amount: number,
        public accountID: string,
        public timestamp: Date,
        public comment?: string,
        public currency: string = 'EUR') { }

    // true, if this expense only exists in local storage and is not yet synced to the server
    public localStorage: Boolean = false;
}
