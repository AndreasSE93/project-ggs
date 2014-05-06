package clientCore;

import java.io.IOException;



import clientHandlers.LobbyHandler;

import clientNetworking.NetManager;
import clientNetworking.Connection;


public class Monitor {
	
	public  Connection conn;
	public  NetManager net;
	//public ChatShell chatModule;
	//public LobbyShell lobbyModule;
	
	public Monitor() {
		
	}
	
	public void init(){
		
		this.conn = new Connection("130.243.137.81", 8080);
		this.net = new NetManager(conn);
        try {
			net.connectToServer();
			setState(1);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
        

      
		
		
	}
	
	public void setState (int state){
		if(state == 1){
			
			LobbyHandler lh = new LobbyHandler(net);
			lh.initHandler();
		}
		else{
			System.out.println("hejdå");
		}
		
		
	}
	
/*	private void runner() {
		String firstcall;
		LobbyServerMessage LM = null;
		try {
			firstcall = net.receiveMessage();
			System.out.println(firstcall + "\n");
			LM = this.lobbyModule.handler.decode(firstcall);
			System.out.println(LM + "\n");
		} catch (IOException | JSONException e) {
			e.printStackTrace();
		}
		if (LM != null) {
			this.lobbyModule.Initiate(LM.UserList, LM.GameHost, this.net);
			this.chatModule.Initiate(this.lobbyModule.graphInterface.chatt, this.net, this.lobbyModule.graphInterface.lobby);
		} else {
			System.out.println("Connection failed");
		}
		
		String mess;
		
		
		while(true) {
			try {
				
				MessageHandler MH = new MessageHandler();
				mess = net.receiveMessage();
				Message m = MH.decode(mess);
				this.chatModule.graphInterface.chatPanel.setText(this.chatModule.graphInterface.chatPanel.getText() + "<" + getTime() + "> " + m.user + ": "  + m.message + "\n");
				this.chatModule.graphInterface.chatPanel.setCaretPosition(this.chatModule.graphInterface.chatPanel.getDocument().getLength());
			} catch (IOException | JSONException e) {
			
				e.printStackTrace();
			}
		}
	}
	*/


	
}