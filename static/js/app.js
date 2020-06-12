//https://stackoverflow.com/questions/31697388/dymo-label-javascript-printing-framework
function startupCode() {
    var printers = dymo.label.framework.getPrinters();
    if (printers.length == 0)
      throw "No DYMO printers are installed. Install DYMO printers.";
  
    dymo.label.framework.getPrintersAsync().then(function(printers){
      // Successful result, printers variable has a list of all supported by the DYMO Label Framework
      console.log(printers);
    }).thenCatch(function(error){
      // Error
    });
  
    var labelXml = 'formatted label here'   
  
    var label = dymo.label.framework.openLabelXml(labelXml);
    label.setObjectText("BARCODE", '000220200');
    label.print("DYMO LabelWriter 450"); // This is the NAME of the printer which i found 
}

function frameworkInitShim() {
    dymo.label.framework.trace = 1; //true
    dymo.label.framework.init(startupCode); 
}

window.onload = frameworkInitShim;