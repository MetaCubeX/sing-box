### 入站

```json
{
  "enabled": true,
  "server_name": "",
  "alpn": [],
  "min_version": "",
  "max_version": "",
  "cipher_suites": [],
  "certificate": "",
  "certificate_path": "",
  "key": "",
  "key_path": "",
  "acme": {
    "domain": [],
    "data_directory": "",
    "default_server_name": "",
    "email": "",
    "provider": "",
    "disable_http_challenge": false,
    "disable_tls_alpn_challenge": false,
    "alternative_http_port": 0,
    "alternative_tls_port": 0,
    "external_account": {
      "key_id": "",
      "mac_key": ""
    }
  }
}
```

### 出站

```json
{
  "enabled": true,
  "disable_sni": false,
  "server_name": "",
  "insecure": false,
  "alpn": [],
  "min_version": "",
  "max_version": "",
  "cipher_suites": [],
  "certificate": "",
  "certificate_path": "",
  "ech": {
    "enabled": false,
    "pq_signature_schemes_enabled": false,
    "dynamic_record_sizing_disabled": false,
    "config": ""
  },
  "utls": {
    "enabled": false,
    "fingerprint": ""
  }
}
```

TLS 版本值：

* `1.0`
* `1.1`
* `1.2`
* `1.3`

密码套件值：

* `TLS_RSA_WITH_AES_128_CBC_SHA`
* `TLS_RSA_WITH_AES_256_CBC_SHA`
* `TLS_RSA_WITH_AES_128_GCM_SHA256`
* `TLS_RSA_WITH_AES_256_GCM_SHA384`
* `TLS_AES_128_GCM_SHA256`
* `TLS_AES_256_GCM_SHA384`
* `TLS_CHACHA20_POLY1305_SHA256`
* `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA`
* `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`
* `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`
* `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`
* `TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`
* `TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`
* `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`
* `TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`
* `TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256`
* `TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`

!!! note ""

    当内容只有一项时，可以忽略 JSON 数组 [] 标签

### 字段

#### enabled

启用 TLS

#### disable_sni

==仅客户端==

不要在 ClientHello 中发送服务器名称.

#### server_name

用于验证返回证书上的主机名，除非设置不安全。

它还包含在 ClientHello 中以支持虚拟主机，除非它是 IP 地址。

#### insecure

==仅客户端==

接受任何服务器证书。

#### alpn

支持的应用层协议协商列表，按优先顺序排列。

如果两个对等点都支持 ALPN，则选择的协议将是此列表中的一个，如果没有相互支持的协议则连接将失败。

参阅 [Application-Layer Protocol Negotiation](https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation)。

#### min_version

可接受的最低 TLS 版本。

默认情况下，当前使用 TLS 1.2 作为客户端的最低要求。作为服务器时使用 TLS 1.0。

#### max_version

可接受的最大 TLS 版本。

默认情况下，当前最高版本为 TLS 1.3。

#### cipher_suites

将在 ECDHE 握手中使用的椭圆曲线，按优先顺序排列。

如果为空，将使用默认值。

客户端将使用第一个首选项作为其在 TLS 1.3 中的密钥共享类型。
这在未来可能会改变。

#### certificate

服务器 PEM 证书。

#### certificate_path

服务器 PEM 证书路径。

#### key

==仅服务器==

服务器 PEM 私钥。

#### key_path

==仅服务器==

服务器 PEM 私钥路径。

#### ech

==仅客户端==

!!! warning ""

    默认安装不包含 ECH, 参阅 [安装](/zh/#_2)。

ECH (Encrypted Client Hello) 是一个 TLS 扩展，它允许客户端加密其 ClientHello 的第一部分
信息。

如果您不知道如何填写其他配置，只需设置 `enabled` 即可。

#### utls

==仅客户端==

!!! warning ""

    默认安装不包含 uTLS, 参阅 [安装](/zh/#_2)。

!!! note ""

    uTLS 维护不善且其效果可能未经证实，使用风险自负。

uTLS 是 "crypto/tls" 的一个分支，它提供了 ClientHello 指纹识别阻力。

可用的指纹值：

* chrome
* firefox
* ios
* android
* random

### ACME 字段

!!! warning ""

    默认安装不包含 ACME，参阅 [安装](/zh/#_2)。

#### domain

一组域名。

默认禁用 ACME。

#### data_directory

ACME 数据目录。

默认使用 `$XDG_DATA_HOME/certmagic|$HOME/.local/share/certmagic`。

#### default_server_name

如果 ClientHello 的 ServerName 字段为空，则选择证书时要使用的服务器名称。

#### email

创建或选择现有 ACME 服务器帐户时使用的电子邮件地址。

#### provider

要使用的 ACME CA 供应商。

| 值                  | 供应商           |
|--------------------|---------------|
| `letsencrypt (默认)` | Let's Encrypt |
| `zerossl`          | ZeroSSL       |
| `https://...`      | 自定义           |

#### disable_http_challenge

禁用所有 HTTP 质询。

#### disable_tls_alpn_challenge

禁用所有 TLS-ALPN 质询。

#### alternative_http_port

用于 ACME HTTP 质询的备用端口；如果非空，将使用此端口而不是 80 来启动 HTTP 质询的侦听器。

#### alternative_tls_port

用于 ACME TLS-ALPN 质询的备用端口； 系统必须将 443 转发到此端口以使质询成功。

#### external_account

EAB（外部帐户绑定）包含将 ACME 帐户绑定或映射到其他已知帐户所需的信息由 CA。

外部帐户绑定“用于将 ACME 帐户与非 ACME 系统中的现有帐户相关联，例如 CA 客户数据库。

为了启用 ACME 帐户绑定，运行 ACME 服务器的 CA 需要向 ACME 客户端提供 MAC 密钥和密钥标识符，使用 ACME 之外的一些机制。
§7.3.4

#### external_account.key_id

密钥标识符。

#### external_account.mac_key

MAC 密钥。

### 重载

对于服务器配置，如果修改，证书和密钥将自动重新加载。