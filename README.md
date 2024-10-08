# data-platform-request-reads-cache-manager-rmq-kube
data-platform-request-reads-cache-manager-rmq-kube は、フロントエンドUI等 からのキャッシュデータ更新要求に基づいて、バックエンドのreads系マイクロサービス郡とのやり取りをコントロールしながら中継仲介する マイクロサービスです。  

## 動作環境
data-platform-request-reads-cache-manager-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## api-input-reader
api-input-reader は、フロントエンドUI等 からのデータ要求入力条件を定義するタイプフォーマットです。  

## api-module-runtimes-requests
api-module-runtimes-requests は、フロントエンドUIからのデータ要求入力条件をもとに、バックエンドに対してリクエストを要求するファンクションプログラムとパターンを記述します。  

## api-module-runtimes-responses
api-module-runtimes-responses は、バックエンドからのデータのレスポンスを定義するタイプフォーマットです。  

## api-output-formatter
api-output-formatter は、レスポンスされたデータをフロントエンドUI等にキャッシュデータとして返す際のフォーマット定義です。

## controllers
controllers は、本マイクロサービスのコアで、UIのキャッシュデータ更新要求仕様を、UIの詳細要求に応じて細かく記述します。  
バックエンドの、しばしば10以上の複数マイクロサービスとの複雑なデータ取得要求が記載されます。　　

## services
services 内の機能は、controller において複数のデータ構造をマッピングして結合したい場合や、全controllerに共通の共通サービスを定義する場合に使用されます。  

## routers
routersは、controller の各機能をUI のURL制御と関連つけるための定義を記述します。  