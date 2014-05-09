package clientCore;

import java.io.IOException;

import javax.swing.JOptionPane;



import clientHandlers.LobbyHandler;
import clientHandlers.TiarHandler;

import clientNetworking.NetManager;
import clientNetworking.Connection;


public class Monitor {
	
	public  Connection conn;
	public  NetManager net;
	public String userName;
	
	public Monitor() {
		
	}
	//130.243.137.247 andreas
	
	public void init(){
		
		this.conn = new Connection("130.243.137.247", 8080);
		this.net = new NetManager(conn);
        try {
			net.connectToServer();
			userName = (String)JOptionPane.showInputDialog("Write username!"); 
			setState(1);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
        

      
		
		
	}
	
	public void setState (int state){
		while(true){
		if(state == 1){
			LobbyHandler lh = new LobbyHandler(net, userName);
			state = lh.init();
			continue;
		}
		if(state == 2){
			 TiarHandler th = new TiarHandler(net, userName);
			 state = th.init();
		}
		else{
			System.out.println("hejdå");
		}
		}
		
	}
	
	
}