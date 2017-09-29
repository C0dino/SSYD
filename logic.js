if (record.totalReaders <= 0) {
    record.readerKit4 = 0;
    record.readerKit8 = 0;
    record.v100Req = 0;
    record.ep1502Req = 0;
    record.mr52Req = 0;
  }
  else if (record.totalReaders > 0 && record.totalReaders <= 4) {
    record.readerKit4 = 1;
    record.readerKit8 = 0;
    record.v100Req = 0;
  } else if (record.totalReaders >= 5 && record.totalReaders <= 8) {
    record.readerKit8 = 1;
    record.readerKit4 = 0;
    record.v100Req = 0;
  } else if (record.totalReaders >=9) {
    record.v100Req = Math.ceil((record.totalReaders-8)/2);
    record.readerKit8 = 1;
    record.readerKit4 = 0;
  }
  
  if (record.totalReaders >=1) {
    record.ep1502Req = Math.ceil(record.totalReaders/64);
    record.mr52Req = Math.ceil(((record.totalReaders - (record.ep1502Req*2))/2));
  }
  
  if (record.totalInputs <= 0) {
    record.v200Req = 0;
    record.mr16INReq = 0;
  }
  if (record.totalOutputs <= 0) {
    record.v300Req = 0;
    record.mr16OUTReqReq = 0;
  }
  
  if (record.totalInputs >= 1 && record.useAux !== false) {
    record.v200Req = Math.ceil((record.totalInputs - ((Math.ceil(record.totalReaders/2)*3)- 3))/16);
    record.mr16INReq = Math.ceil((record.totalInputs - ((Math.ceil(record.totalReaders/2)*2)/16)));
  } else if (record.totalInputs >= 1 && record.useAux !== true) {
    record.v200Req = Math.ceil(record.totalInputs/16);
    record.mr16INReq = Math.ceil(record.totalInputs/16);
  }
  if (record.totalOutputs >= 1 && record.useAuxO !== false) {
    record.v300Req = Math.ceil((record.totalOutputs - ((Math.ceil(record.totalReaders/2)*2)+(record.v200Req * 2)))/12);
    record.mr16OUTReq = Math.ceil((record.totalOutputs - ((Math.ceil(record.totalReaders/2)*2)+(record.mr16INReq * 2)))/16);
  } else if (record.totalOutputs >= 1 && record.useAuxO !== true) {
    record.v300Req = Math.ceil(record.totalOutputs/12);
    record.mr16OUTReq = Math.ceil(record.totalOutputs/16);
  }
  
  if (record.v100Req >= 1 || record.v200Req >= 1 || record.v300Req >= 1) {
    record.boardsRemaining =  (((record.v100Req + record.v200Req + record.v300Req)/5) - Math.floor((record.v100Req + record.v200Req + record.v300Req)/5));
    if ((record.boardsRemaining <= 0) || (record.boardsRemaining >= 0.7)){
    record.enclosureL = Math.ceil((record.v100Req + record.v200Req + record.v300Req)/5);
    record.enclosureS = 0;
    } else if (record.boardsRemaining <= 0.6 && record.boardsRemaining > 0) {
    record.enclosureL = Math.floor((record.v100Req + record.v200Req + record.v300Req)/5);
    record.enclosureS = 1;
    }
  if (record.enclosureL >= 1 || record.enclosureS >= 1) {
   record.enclosurePSU = (record.enclosureL + record.enclosureS); 
  }
  } if (record.v100Req <= 0 && record.v200Req <= 0 && record.v300Req <= 0) {
    record.enclosureS = 0;
    record.enclosureL = 0;
    record.enclosurePSU = 0;
  }