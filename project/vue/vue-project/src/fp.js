function addScript(url){
	document.write("<script language=javascript src="+url+"></script>");
}
addScript("https://cdn.bootcss.com/fingerprintjs2/2.1.0/fingerprint2.min.js");
function createFingerprint() {
  // 浏览器指纹
  const fingerprint = Fingerprint2.get({
    excludes: { // 排除的配置信息
      plugins: true, // 排除浏览器插件
      canvas : true,
      webgl : true,
      touchSupport : true,
  }} , (components) => { // 参数只有回调函数时，默认浏览器指纹依据所有配置信息进行生成
    const values = components.map(component => component.value); // 配置的值的数组
    const murmur = Fingerprint2.x64hash128(values.join(''), 31); // 生成浏览器指纹
    console.log(murmur);
    // components.fp = murmur;
    console.log(components);
    // 这里的 components 是一个列表 
     // 这里是一个list
    let data = {};
    for (let i = 0; i < components.length; i++) {
      const element = components[i]; 
      // console.log(element.key + ' : ' + element.value);
      data[element.key] = element.value; 
    }
    data.devicehashvalue = murmur;
    console.log(data);
    // TODO 将其发送到我们的后台
    sendFingerprint('http://localhost:3000/fp',data);
    localStorage.setItem('browserId', murmur); // 存储浏览器指纹，在项目中用于校验用户身份和埋点
  });
}
// 发送指纹到指定服务器
function sendFingerprint(url , data){
  var xhr = new XMLHttpRequest();
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.setRequestHeader('Access-Control-Allow-Credentials', 'true');
  xhr.setRequestHeader('Access-Control-Allow-Origin', '*');
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
      console.log(xhr.responseText);
    }
  };
  xhr.send(JSON.stringify(data));
}

function GetFinger() {
  // 您不应在页面加载时或加载后直接运行指纹。 而是使用setTimeout或requestIdleCallback将其延迟几毫秒，以确保指纹一致。
  if (window.requestIdleCallback) {
    requestIdleCallback(() => {
      this.createFingerprint();
    });
  } else {
    setTimeout(() => {
      this.createFingerprint();
    }, 500);
  }
}

GetFinger();






// 理想化 
// http 进行发送 